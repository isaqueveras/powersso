// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postgres

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/oops"
)

type PGOTP struct {
	DB     *postgres.Transaction
	UserID *uuid.UUID
}

func (pg *PGOTP) GetToken() (userName, token *string, err error) {
	if err = pg.DB.Builder.
		Select("CONCAT('(',first_name,' ',last_name,')'), otp").
		From("public.users").
		Where("id = ?::UUID AND otp NOTNULL", pg.UserID).
		QueryRow().
		Scan(&userName, &token); err != nil {
		return nil, nil, oops.Err(err)
	}

	return
}

func (pg *PGOTP) SetToken(secret *string) (err error) {
	if _, err = pg.DB.Builder.
		Update("users").
		Set("otp", secret).
		Where("id = ?", pg.UserID).
		Exec(); err != nil {
		return oops.Err(err)
	}
	return
}
