// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package roles

import (
	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
)

var _ roles.IRoles = (*repository)(nil)

// repository is the implementation of the roles repository
type repository struct {
	pg *pgRoles
}

// New creates a new repository
func New(transaction *postgres.DBTransaction) roles.IRoles {
	return &repository{pg: &pgRoles{DB: transaction}}
}

// RemoveRoles contains the flow for the remove roles
func (r *repository) RemoveRoles(userID *string, roles string) error {
	return r.pg.removeRoles(userID, roles)
}
