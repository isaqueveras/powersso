// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// PGOTP is the implementation of transaction for the otp repository
type PGOTP struct {
	DB *postgres.DBTransaction
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
