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

var _ domain.IRole = (*repoRole)(nil)

type repoRole struct {
	pg *infra.PGRole
}

// NewRoleRepository creates a new repository
func NewRoleRepository(tx *pg.Transaction) domain.IRole {
	return &repoRole{pg: &infra.PGRole{DB: tx}}
}

// Set put the flag value in the database
func (r *repoRole) Set(userID *uuid.UUID, flag *domain.Flag) error {
	return r.pg.Set(userID, flag)
}
