// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package session

import (
	"github.com/isaqueveras/power-sso/internal/domain/session"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/pkg/query"
)

// pgSession is the implementation
// of transaction for the session repository
type pgSession struct {
	DB *postgres.DBTransaction
}

// create add session of the user in database
func (pg *pgSession) create(in *session.Session) (err error) {
	_cols, _vals, err := query.FormatValuesInUp(in)
	if err != nil {
		return oops.Err(err)
	}

	if err = pg.DB.Builder.
		Insert("sessions").
		Columns(_cols...).
		Values(_vals...).
		Suffix(`RETURNING "id"`).
		Scan(new(string)); err != nil {
		return oops.Err(err)
	}

	return
}
