package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/oops"
)

type PGUser struct {
	DB *postgres.Transaction
}

func (pg *PGUser) Exist(email *string) (err error) {
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

func (pg *PGUser) Get(data *domain.User) (err error) {
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

func (pg *PGUser) Disable(userUUID *uuid.UUID) (err error) {
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

func (pg *PGUser) ChangePassword(in *domain.ChangePassword) error {
	if err := pg.DB.Builder.
		Update("users").
		Set("password", in.Password).
		Set("attempts", 0).
		Set("last_failure", squirrel.Expr("NULL")).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{"id": in.UserID, "active": true}).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return nil
}
