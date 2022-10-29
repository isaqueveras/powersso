// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"context"

	domain "github.com/isaqueveras/power-sso/internal/domain/project"
	infra "github.com/isaqueveras/power-sso/internal/infrastructure/project"
	"github.com/isaqueveras/power-sso/pkg/conversor"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// Create is the business logic for creating a project
func Create(ctx context.Context, input *CreateProjectReq) (err error) {
	var (
		transaction *postgres.DBTransaction
		data        *domain.CreateProject
	)

	if transaction, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	if data, err = conversor.TypeConverter[domain.CreateProject](&input); err != nil {
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
