package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	domain "github.com/isaqueveras/powersso/internal/domain/auth"
	pg "github.com/isaqueveras/powersso/pkg/database/postgres"
	"github.com/isaqueveras/powersso/pkg/oops"
)

// PGUser is the implementation of transaction for the user repository
type PGUser struct {
	DB *pg.Transaction
}

// Exist check if the user exists by email in the database
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

// Get get the user from the database
func (pg *PGUser) Get(data *domain.User) (err error) {
	if err = pg.DB.Builder.
		Select(`id, email, first_name, last_name, about, key, active, level, otp`).
		Column("attempts >= 3 AND (last_failure + '5 minutes') >= NOW() AS blocked").
		Column("(flag & ?) <> 0", domain.FlagOTPEnable).
		Column("(flag & ?) <> 0", domain.FlagOTPSetup).
		From("users").
		Where(squirrel.Or{squirrel.Eq{"id": data.ID}, squirrel.Eq{"email": data.Email}}).
		Scan(&data.ID, &data.Email, &data.FirstName, &data.LastName, &data.About, &data.Key,
			&data.Active, &data.Level, &data.OTPToken, &data.Blocked, &data.OTPEnable, &data.OTPSetUp); err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrUserNotExists()
		}
		return oops.Err(err)
	}

	return
}

// Disable disable user in database
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
