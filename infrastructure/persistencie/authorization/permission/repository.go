// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package permission

import (
	"context"

	"github.com/google/uuid"
	database "github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/authorization/permission"
	"github.com/isaqueveras/powersso/infrastructure/persistencie/authorization/permission/postgres"
)

var _ permission.IPermission = (*repository)(nil)

type repository struct {
	ctx context.Context
	pg  *postgres.Database
}

// NewRepository ...
func NewRepository(ctx context.Context, tx *database.Transaction) permission.IPermission {
	return &repository{pg: postgres.New(tx)}
}

// Get ...
func (r *repository) Get(userID, organizationID *uuid.UUID) ([]*string, error) {
	return r.pg.Get(r.ctx, userID, organizationID)
}

// Create ...
func (r *repository) Create(in *permission.Permission) error {
	return r.pg.Create(r.ctx, in)
}
