// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"context"

	"github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/project"
)

// database is the implementation of transaction for the session repository
type database struct {
	DB *postgres.Transaction
}

// createNewProject contains the flow for create a project in database
func (db *database) createNewProject(ctx context.Context, in *project.CreateProject) error {
	_, err := db.DB.Builder.
		Insert("projects").
		Columns("name", "desc", "slug", "created_by", "url").
		Values(in.Name, in.Desc, in.Slug, in.CreatedByID, in.Url).
		ExecContext(ctx)
	return err
}
