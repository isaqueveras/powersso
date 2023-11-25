// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"context"

	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/project"
)

// repository is the implementation of the session repository
type repository struct {
	pg  *database
	ctx context.Context
}

// New creates a new repository
func New(ctx context.Context, tx *postgres.Transaction) project.IProject {
	return &repository{ctx: ctx, pg: &database{DB: tx}}
}

// CreateNewProject contains the flow for create project in database
func (r *repository) CreateNewProject(input *project.CreateProject) error {
	return r.pg.createNewProject(r.ctx, input)
}
