// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package session

import (
	"github.com/isaqueveras/power-sso/internal/domain/session"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
)

// repository is the implementation of the session repository
type repository struct {
	pg *pgSession
}

// New creates a new repository
func New(transaction *postgres.DBTransaction) session.ISession {
	return &repository{
		pg: &pgSession{DB: transaction},
	}
}

// Create contains the flow for create session in database
func (r *repository) Create(userID, clientIP, userAgent *string) (*string, error) {
	return r.pg.create(userID, clientIP, userAgent)
}

// Delete contains the flow for Delete user
func (r *repository) Delete(sessionID *string) error {
	return r.pg.delete(sessionID)
}
