package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	database "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/authentication"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

type (
	// PGAuth is the implementation of transaction for the auth repository
	PGAuth struct{ DB *database.Transaction }

	// PGOTP is the implementation of transaction for the user repository
	User struct{ DB *database.Transaction }

	// OTP is the implementation of transaction for the otp repository
	OTP struct{ DB *database.Transaction }

	// Session is the implementation of transaction for the session repository
	Session struct{ DB *database.Transaction }

	// Flag is the implementation of transaction for the flag repository
	Flag struct{ DB *database.Transaction }
)

// CreateAccount register the user in the database
func (pg *PGAuth) CreateAccount(input *domain.CreateAccount) (userID *uuid.UUID, err error) {
	_cols, _vals, err := utils.FormatValuesInUp(input)
	if err != nil {
		return nil, oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("public.user").
		Columns(_cols...).
		Values(_vals...).
		Suffix(`RETURNING "id"`).
		Scan(&userID); err != nil {
		return nil, oops.Err(err)
	}

	return
}

func (pg *PGAuth) AddAttempts(userID *uuid.UUID) (err error) {
	if _, err = pg.DB.Builder.
		Update("public.user").
		Set("attempts", squirrel.Expr("attempts + 1")).
		Set("last_failure", squirrel.Expr("NOW()")).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	return
}

func (pg *PGAuth) LoginSteps(email *string) (steps *domain.Steps, err error) {
	steps = new(domain.Steps)
	if err = pg.DB.Builder.
		Select("first_name").
		Column("(flag&?) <> 0 AND 	(flag&?) <> 0",
			domain.FlagOTPEnable, domain.FlagOTPSetup).
		From("public.user").
		Where("email = ?", email).
		Limit(1).
		Scan(&steps.Name, &steps.OTP); err != nil && err != sql.ErrNoRows {
		return nil, oops.Err(err)
	}

	return
}

// AccountExists validate whether an account with the same identifier already exists
func (pg *User) AccountExists(email *string) (err error) {
	var exists *bool
	if err = pg.DB.Builder.
		Select("COUNT(id) > 0").
		From("public.user").
		Where(squirrel.Eq{"email": email}).
		Scan(&exists); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	if exists != nil && *exists {
		return domain.ErrUserExists()
	}

	return
}

// GetUser fetches a user's data from the database
func (pg *User) GetUser(data *domain.User) (err error) {
	cond := squirrel.Eq{"id": data.ID}
	if data.Email != nil {
		cond = squirrel.Eq{"email": data.Email}
	}

	if err = pg.DB.Builder.
		Select(`id, email, password, first_name, last_name, flag, key, active, level, otp`).
		Column("attempts >= 3 AND (last_failure + '5 minutes') >= NOW() AS blocked").
		Column("(flag & ?) <> 0", domain.FlagOTPEnable).
		Column("(flag & ?) <> 0", domain.FlagOTPSetup).
		From("public.user").
		Where(cond).
		Scan(&data.ID, &data.Email, &data.Password, &data.FirstName, &data.LastName, &data.Flag, &data.Key,
			&data.Active, &data.Level, &data.OTPToken, &data.Blocked, &data.OTPEnable, &data.OTPSetUp); err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrUserNotExists()
		}
		return oops.Err(err)
	}

	return
}

func (pg *User) Disable(userUUID *uuid.UUID) (err error) {
	if err = pg.DB.Builder.
		Update("public.user").
		Set("active", false).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{"id": userUUID, "active": true}).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return
}

func (pg *User) ChangePassword(in *domain.ChangePassword) error {
	if err := pg.DB.Builder.
		Update("public.user").
		Set("password", in.Password).
		Set("attempts", 0).
		Set("key", in.Key).
		Set("last_failure", squirrel.Expr("NULL")).
		Where(squirrel.Eq{"id": in.UserID, "active": true}).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return nil
}

func (pg *OTP) GetToken(userID *uuid.UUID) (userName, token *string, err error) {
	if err = pg.DB.Builder.
		Select("CONCAT('(',first_name,' ',last_name,')'), otp").
		From("public.user").
		Where("id = ?::UUID AND otp NOTNULL", userID).
		QueryRow().
		Scan(&userName, &token); err != nil {
		return nil, nil, oops.Err(err)
	}

	return
}

func (pg *OTP) SetToken(userID *uuid.UUID, secret *string) (err error) {
	if _, err = pg.DB.Builder.
		Update("public.user").
		Set("otp", secret).
		Where("id = ?", userID).
		Exec(); err != nil {
		return oops.Err(err)
	}
	return
}

// Create add session of the user in database
func (pg *Session) Create(userID *uuid.UUID, clientIP, userAgent *string) (sessionID *uuid.UUID, err error) {
	if err = pg.DB.Builder.
		Insert("session").
		Columns("user_id", "expires_at", "ip", "user_agent").
		Values(userID, squirrel.Expr("NOW() + '15 minutes'"), clientIP, userAgent).
		Suffix(`RETURNING "id"`).
		Scan(&sessionID); err != nil {
		return nil, oops.Err(err)
	}

	if _, err = pg.DB.Builder.
		Update("session").
		Set("deleted_at", squirrel.Expr("NOW()")).
		Where(`id NOT IN (
			SELECT id FROM session
			WHERE user_id = ? AND deleted_at IS NULL
			ORDER BY created_at DESC 
			LIMIT ?
		)`, userID, config.Get().Server.OpenSessionsPerUser).
		Where("user_id = ?", userID).
		Exec(); err != nil {
		return nil, oops.Err(err)
	}

	if _, err = pg.DB.Builder.
		Update("public.user").
		Set("attempts", 0).
		Set("last_login", squirrel.Expr("NOW()")).
		Set("last_failure", nil).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return nil, oops.Err(err)
	}

	return
}

// Delete delete session of the user in database
func (pg *Session) Delete(ids ...*uuid.UUID) (err error) {
	if _, err = pg.DB.Builder.
		Update("session").
		Set("deleted_at", squirrel.Expr("NOW()")).
		Where("deleted_at IS NULL").
		Where(squirrel.Eq{"id": ids}).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	return
}

func (pg *Session) Get(userID *uuid.UUID) (sessions []*uuid.UUID, err error) {
	query := pg.DB.Builder.Select("id").From("session").Where("user_id = ? AND deleted_at IS NULL", userID)

	row, err := query.Query()
	if err != nil {
		return nil, oops.Err(err)
	}

	for row.Next() {
		var sessionID *uuid.UUID
		if err = row.Scan(&sessionID); err != nil {
			return nil, oops.Err(err)
		}
		sessions = append(sessions, sessionID)
	}

	return
}

func (pg *Flag) Set(userID *uuid.UUID, flag domain.Flag) error {
	if _, err := pg.DB.Builder.
		Update("public.user").
		Set("flag", flag).
		Where("id = ?::UUID", userID).
		Exec(); err != nil {
		return oops.Err(err)
	}
	return nil
}

func (pg *Flag) Get(userID *uuid.UUID) (flag *int64, err error) {
	if err = pg.DB.Builder.
		Select("flag").
		From("public.user").
		Where("id = ?::UUID", userID).
		Scan(&flag); err != nil {
		return nil, oops.Err(err)
	}
	return
}
