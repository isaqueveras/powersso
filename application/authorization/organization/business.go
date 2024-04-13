// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package organization

import (
	"context"

	db "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/authorization/organization"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/authorization/organization"
	"github.com/isaqueveras/powersso/oops"
)

// Create is the business logic for creating a new room
func Create(ctx context.Context, in *Organization) (err error) {
	var tx *db.Transaction
	if tx, err = db.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repo := infra.New(ctx, tx)
	if err = repo.Create(&domain.Organization{
		Name:        in.Name,
		Desc:        in.Description,
		URL:         in.Url,
		CreatedByID: in.CreatedByID,
	}); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}
