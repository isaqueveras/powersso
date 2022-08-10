// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"database/sql"

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
func (pg *pgAuth) createAccessToken(userID *string) (err error) {
	if _, err = pg.DB.Execute(`
		INSERT INTO activate_account_tokens (user_id, expires_at) 
		VALUES ($1, now() + interval '15 minutes')`, userID); err != nil {
		return oops.Err(err)
	}

	return nil
}

// getActivateAccountToken get the activate account token from the database
func (pg *pgAuth) getActivateAccountToken(token *string) (res *auth.ActivateAccountToken, err error) {
	res = new(auth.ActivateAccountToken)

	err = pg.DB.Builder.
		Select(`
			AAT.id,
			AAT.token,
			AAT.user_id,
			AAT.used,
			AAT.expires_at >= now() AS "valid",
			AAT.expires_at,
			AAT.created_at,
			AAT.updated_at`).
		From("activate_account_tokens AAT").
		Where("AAT.token = ?", token).
		Limit(1).
		Scan(&res.ID, &res.Token, &res.UserID, &res.Used, &res.IsValid,
			&res.ExpiresAt, &res.CreatedAt, &res.UpdatedAt)

	if err != nil && err != sql.ErrNoRows {
		return nil, oops.Err(err)
	}

	return
}
