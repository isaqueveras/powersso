// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"context"

	db "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/project"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/project"
	"github.com/isaqueveras/powersso/oops"
)

// CreateNewProject is the business logic for creating a new project
func CreateNewProject(ctx context.Context, in *NewProject) (err error) {
	var tx *db.Transaction
	if tx, err = db.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	data := &domain.CreateProject{
		Name:        in.Name,
		Desc:        in.Description,
		Slug:        in.Slug,
		Url:         in.Url,
		CreatedByID: in.CreatedByID,
	}

	repo := infra.New(ctx, tx)
	if err = repo.CreateNewProject(data); err != nil {
		return oops.Err(err)
	}

	return tx.Commit()
}
