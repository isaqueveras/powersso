// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/pkg/query"
)

type pgAuth struct {
	DB *postgres.DBTransaction
}

func (pg *pgAuth) register(input *auth.Register) (err error) {
	_cols, _vals, err := query.FormatValuesInUp(input)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("users").
		Columns(_cols...).
		Values(_vals...).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return
}
