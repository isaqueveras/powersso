// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package session

import (
	"github.com/Masterminds/squirrel"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// pgSession is the implementation
// of transaction for the session repository
type pgSession struct {
	DB *postgres.DBTransaction
}

// create add session of the user in database
func (pg *pgSession) create(userID *string) (sessionID *string, err error) {
	if err = pg.DB.Builder.
		Insert("sessions").
		Columns("user_id", "expires_at").
		Values(userID, squirrel.Expr("NOW() + '15 minutes'")).
		Suffix(`RETURNING "id"`).
		Scan(&sessionID); err != nil {
		return nil, oops.Err(err)
	}

	return
}
