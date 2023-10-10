package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	database "github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/auth"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

type (
	// PGAuth is the implementation of transaction for the auth repository
	PGAuth struct{ DB *database.Transaction }

	// PGOTP is the implementation of transaction for the user repository
	User struct{ DB *database.Transaction }

	// A2F is the implementation of transaction for the otp repository
	A2F struct{ DB *database.Transaction }

	// Session is the implementation of transaction for the session repository
	Session struct{ DB *database.Transaction }

	// Flag is the implementation of transaction for the flag repository
	Flag struct{ DB *database.Transaction }
)

// CreateAccount register the user in the database
func (pg *PGAuth) CreateAccount(input *auth.CreateAccount) (userID *uuid.UUID, err error) {
	_cols, _vals, err := utils.FormatValuesInUp(input)
	if err != nil {
		return nil, oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("users").
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
		Update("users").
		Set("attempts", squirrel.Expr("attempts + 1")).
		Set("last_failure", squirrel.Expr("NOW()")).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	return
}

func (pg *PGAuth) LoginSteps(email *string) (steps *auth.Steps, err error) {
	steps = new(auth.Steps)
	if err = pg.DB.Builder.
		Select("first_name").
		Column("(flag&?) <> 0 AND 	(flag&?) <> 0",
			auth.FlagOTPEnable, auth.FlagOTPSetup).
		From("users").
		Where("email = ?", email).
		Limit(1).
		Scan(&steps.Name, &steps.OTP); err != nil && err != sql.ErrNoRows {
		return nil, oops.Err(err)
	}

	return
}

func (pg *User) Exist(email *string) (err error) {
	var exists *bool
	if err = pg.DB.Builder.
		Select("COUNT(id) > 0").
		From("users").
		Where(squirrel.Eq{"email": email}).
		Scan(&exists); err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrUserNotExists()
		}
		return oops.Err(err)
	}

	if exists != nil && *exists {
		return domain.ErrUserExists()
	}

	return
}

func (pg *User) Get(data *domain.User) (err error) {
	cond := squirrel.Eq{"id": data.ID}
	if data.Email != nil {
		cond = squirrel.Eq{"email": data.Email}
	}

	if err = pg.DB.Builder.
		Select(`id, email, password, first_name, last_name, flag, key, active, level, otp`).
		Column("attempts >= 3 AND (last_failure + '5 minutes') >= NOW() AS blocked").
		Column("(flag & ?) <> 0", domain.FlagOTPEnable).
		Column("(flag & ?) <> 0", domain.FlagOTPSetup).
		From("users").
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
		Update("users").
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
		Update("users").
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

func (pg *A2F) GetToken(userID *uuid.UUID) (userName, token *string, err error) {
	if err = pg.DB.Builder.
		Select("CONCAT('(',first_name,' ',last_name,')'), otp").
		From("public.users").
		Where("id = ?::UUID AND otp NOTNULL", userID).
		QueryRow().
		Scan(&userName, &token); err != nil {
		return nil, nil, oops.Err(err)
	}

	return
}

func (pg *A2F) SetToken(userID *uuid.UUID, secret *string) (err error) {
	if _, err = pg.DB.Builder.
		Update("users").
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
		Insert("sessions").
		Columns("user_id", "expires_at", "ip", "user_agent").
		Values(userID, squirrel.Expr("NOW() + '15 minutes'"), clientIP, userAgent).
		Suffix(`RETURNING "id"`).
		Scan(&sessionID); err != nil {
		return nil, oops.Err(err)
	}

	if _, err = pg.DB.Builder.
		Update("users").
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
		Update("sessions").
		Set("deleted_at", squirrel.Expr("NOW()")).
		Where("deleted_at IS NULL").
		Where(squirrel.Eq{"id": ids}).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	return
}

func (pg *Session) Get(userID *uuid.UUID) (sessions []*uuid.UUID, err error) {
	query := pg.DB.Builder.Select("id").From("sessions").Where("user_id = ? AND deleted_at IS NULL", userID)

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

func (pg *Flag) Set(userID *uuid.UUID, flag auth.Flag) error {
	if _, err := pg.DB.Builder.
		Update("users").
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
		From("users").
		Where("id = ?::UUID", userID).
		Scan(&flag); err != nil {
		return nil, oops.Err(err)
	}
	return
}
