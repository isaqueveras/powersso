// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	domain "github.com/isaqueveras/powersso/internal/domain/auth"
	infra "github.com/isaqueveras/powersso/internal/infrastructure/persistencie/auth/postgres"
	pg "github.com/isaqueveras/powersso/pkg/database/postgres"
)

var _ domain.ISession = (*repoSession)(nil)

type repoSession struct{ pg *infra.PGSession }

// NewSessionRepository creates a new repository
func NewSessionRepository(tx *pg.Transaction) domain.ISession {
	return &repoSession{pg: &infra.PGSession{DB: tx}}
}

// Create create a new session for a user
func (r *repoSession) Create(userID *uuid.UUID, clientIP, userAgent *string) (*uuid.UUID, error) {
	return r.pg.Create(userID, clientIP, userAgent)
}

// Delete delete a session for a user
func (r *repoSession) Delete(sessionID *uuid.UUID) error {
	return r.pg.Delete(sessionID)
}
