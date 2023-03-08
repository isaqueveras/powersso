// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/internal/domain/auth/user"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// pgUser is the implementation
// of transaction for the user repository
type pgUser struct {
	DB *postgres.DBTransaction
}

// findByEmailUserExists check if the user exists by email in the database
func (pg *pgUser) findByEmailUserExists(email *string) (exists bool, err error) {
	if err = pg.DB.Builder.
		Select("COUNT(id) > 0").
		From("users").
		Where(squirrel.Eq{
			"email": email,
		}).
		Scan(&exists); err != nil && err != sql.ErrNoRows {
		return false, oops.Err(err)
	}

	return
}

// getUser get the user from the database
func (pg *pgUser) getUser(data *user.User) (err error) {
	if err = pg.DB.Builder.
		Select(`
			U.id,
			U.email,
			U.first_name,
			U.last_name,
			U.roles,
			U.about,
			U.avatar,
			U.phone_number,
			U.address,
			U.city,
			U.country,
			U.gender,
			U.postcode,
			U.token_key,
			U.birthday,
			U.created_at,
			U.updated_at,
			U.login_date,
			U.is_active,
			U.user_type,
			U.number_failed_attempts >= 3 AND (U.last_failure_date + '1 hour') >= NOW() AS blocked_temporarily,
			U.otp,
			U.otp_token,
			U.otp_setup`).
		From("users U").
		Where(squirrel.Or{
			squirrel.Eq{"id": data.ID},
			squirrel.Eq{"email": data.Email},
		}).
		Scan(&data.ID, &data.Email, &data.FirstName, &data.LastName, &data.Roles,
			&data.About, &data.Avatar, &data.PhoneNumber, &data.Address, &data.City,
			&data.Country, &data.Gender, &data.Postcode, &data.TokenKey, &data.Birthday,
			&data.CreatedAt, &data.UpdatedAt, &data.LoginDate, &data.IsActive,
			&data.UserType, &data.BlockedTemporarily, &data.OTPEnabled,
			&data.OTPToken, &data.OTPSetup); err != nil {
		return oops.Err(err)
	}

	return
}

// disableUser disable user in database
func (pg *pgUser) disableUser(userUUID *uuid.UUID) (err error) {
	if err = pg.DB.Builder.
		Update("users").
		Set("is_active", false).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			"id":        userUUID,
			"is_active": true,
		}).
		Suffix("RETURNING id").
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return
}
