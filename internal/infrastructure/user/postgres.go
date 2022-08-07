// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"database/sql"

	"github.com/Masterminds/squirrel"

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
