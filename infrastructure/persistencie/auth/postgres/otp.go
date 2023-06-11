// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	pg "github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/oops"
)

// PGOTP is the implementation of transaction for the otp repository
type PGOTP struct {
	DB *pg.Transaction
}

// GetToken fetch the token of a user's otp
func (pg *PGOTP) GetToken(userID *uuid.UUID) (userName, token *string, err error) {
	if err = pg.DB.Builder.
		Select("CONCAT('(',first_name,' ',last_name,')') AS user_name, otp_token").
		From("public.users").
		Where("id = ?::UUID", userID).
		QueryRow().
		Scan(&userName, &token); err != nil {
		return nil, nil, oops.Err(err)
	}
	return
}

// Configure configure otp for a user
func (pg *PGOTP) Configure(userID *uuid.UUID, secret *string) (err error) {
	if _, err = pg.DB.Builder.
		Update("users").
		Set("otp_token", secret).
		Set("otp", true).
		Set("otp_setup", true).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}
	return
}

// Unconfigure unconfigure otp for a user
func (pg *PGOTP) Unconfigure(userID *uuid.UUID) (err error) {
	if _, err = pg.DB.Builder.
		Update("users").
		Set("otp", false).
		Set("updated_at", squirrel.Expr("NOW()")).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}
	return
}
