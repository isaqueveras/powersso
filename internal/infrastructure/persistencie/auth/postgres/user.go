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
		Select(`
			id,
			email,
			first_name,
			last_name,
			about,
			avatar,
			token_key,
			created_at,
			last_login,
			active,
			user_type,
			attempts >= 3 AND (U.last_failure + '30 minutes') >= NOW() AS blocked,
			otp,
			otp_token,
			otp_setup`).
		From("users").
		Where(squirrel.Or{squirrel.Eq{"id": data.ID}, squirrel.Eq{"email": data.Email}}).
		Scan(
			&data.ID,
			&data.Email,
			&data.FirstName,
			&data.LastName,
			&data.About,
			&data.Avatar,
			&data.Key,
			&data.CreatedAt,
			&data.LastLogin,
			&data.Active,
			&data.Level,
			&data.Blocked,
			&data.OTPToken,
		); err != nil {
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
