// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"context"

	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/project"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/project"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// Create is the business logic for creating a project
func Create(ctx context.Context, input *CreateProjectReq) (err error) {
	var transaction *postgres.Transaction
	if transaction, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	var data *domain.CreateProject
	if data, err = utils.TypeConverter[domain.CreateProject](&input); err != nil {
		return oops.Err(err)
	}

	repo := infra.New(transaction)
	if err = repo.Create(data); err != nil {
		return oops.Err(err)
	}

	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}
