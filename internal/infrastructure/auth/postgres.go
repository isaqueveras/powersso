// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"

	"github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/pkg/query"
)

// pgAuth is the implementation
// of transaction for the auth repository
type pgAuth struct {
	DB *postgres.DBTransaction
}

// register register the user in the database
func (pg *pgAuth) register(input *auth.Register) (userID *string, err error) {
	_cols, _vals, err := query.FormatValuesInUp(input)
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

// createAccessToken create the access token for the user
func (pg *pgAuth) createAccessToken(userID *string) (token string, err error) {
	if err = pg.DB.Builder.
		Insert("activate_account_tokens").
		Columns("user_id", "expires_at").
		Values(userID, time.Now().Add(15*time.Minute)).
		Suffix(`RETURNING "id"`).
		Scan(&token); err != nil {
		return token, oops.Err(err)
	}

	return
}

// getActivateAccountToken get the activate account token from the database
func (pg *pgAuth) getActivateAccountToken(token *string) (res *auth.ActivateAccountToken, err error) {
	res = new(auth.ActivateAccountToken)

	err = pg.DB.Builder.
		Select(`
			id,
			user_id,
			used,
			expires_at >= now() AS "valid",
			expires_at,
			created_at,
			updated_at`).
		From("activate_account_tokens").
		Where("id = ?", token).
		Limit(1).
		Scan(&res.ID, &res.UserID, &res.Used, &res.IsValid,
			&res.ExpiresAt, &res.CreatedAt, &res.UpdatedAt)

	if err != nil && err != sql.ErrNoRows {
		return nil, oops.Err(err)
	}

	return
}

// markTokenAsUsed mark the token as used in the database
func (pg *pgAuth) markTokenAsUsed(token *string) (err error) {
	if _, err = pg.DB.Builder.
		Update("activate_account_tokens").
		Set("used", true).
		Set("updated_at", time.Now()).
		Where("id = ?", token).
		Exec(); err != nil {
		return oops.Err(err)
	}

	return
}

// login get the user password from the database
func (pg *pgAuth) login(email *string) (password *string, err error) {
	if err = pg.DB.Builder.
		Select("password").
		From("users").
		Where("email = ?", email).
		Limit(1).
		Scan(&password); err != nil {
		return nil, oops.Err(err)
	}

	return
}

func (pg *pgAuth) addNumberFailedAttempts(userID *string) (err error) {
	if _, err = pg.DB.Builder.
		Update("users").
		Set("number_failed_attempts", squirrel.Expr("number_failed_attempts + 1")).
		Set("last_failure_date", squirrel.Expr("NOW()")).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	return
}
