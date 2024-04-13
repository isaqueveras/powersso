// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package organization

import (
	"context"

	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/authorization/organization"
	"github.com/isaqueveras/powersso/oops"
)

// repository is the implementation of the session repository
type repository struct {
	pg  *postgres.Transaction
	ctx context.Context
}

// New creates a new repository
func New(ctx context.Context, tx *postgres.Transaction) organization.IOrganization {
	return &repository{ctx: ctx, pg: tx}
}

// Create contains the flow for create organization in database
func (r *repository) Create(org *organization.Organization) error {
	if _, err := r.pg.Builder.
		Insert("organization").
		Columns("name", "desc", "created_by", "url").
		Values(org.Name, org.Desc, org.CreatedByID, org.URL).
		ExecContext(r.ctx); err != nil {
		return oops.Err(err)
	}
	return nil
}
