// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postgres

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/pkg/database/postgres"
	"github.com/isaqueveras/powersso/pkg/oops"
)

// PGSession is the implementation of transaction for the session repository
type PGSession struct {
	DB *postgres.Transaction
}

// Create add session of the user in database
func (pg *PGSession) Create(userID *uuid.UUID, clientIP, userAgent *string) (sessionID *uuid.UUID, err error) {
	if err = pg.DB.Builder.
		Insert("sessions").
		Columns("user_id", "expires_at", "ip", "user_agent").
		Values(userID, squirrel.Expr("NOW() + '15 minutes'"), clientIP, userAgent).
		Suffix(`RETURNING "id"`).
		Scan(&sessionID); err != nil {
		return nil, oops.Err(err)
	}

	if _, err = pg.DB.Builder.
		Update("users").
		Set("attempts", 0).
		Set("last_login", squirrel.Expr("NOW()")).
		Set("last_failure", nil).
		Where("id = ?", userID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return nil, oops.Err(err)
	}

	return
}

// Delete delete session of the user in database
func (pg *PGSession) Delete(sessionID *uuid.UUID) (err error) {
	if _, err = pg.DB.Builder.
		Update("sessions").
		Set("deleted_at", squirrel.Expr("NOW()")).
		Where("id = ? AND deleted_at IS NULL", sessionID).
		Exec(); err != nil && err != sql.ErrNoRows {
		return oops.Err(err)
	}

	return
}
