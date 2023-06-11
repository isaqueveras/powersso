// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/project"
)

// repository is the implementation of the session repository
type repository struct {
	pg *pg
}

// New creates a new repository
func New(transaction *postgres.Transaction) project.IProject {
	return &repository{pg: &pg{DB: transaction}}
}

// Create contains the flow for create project in database
func (r *repository) Create(input *project.CreateProject) error {
	return r.pg.create(input)
}
