// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/isaqueveras/power-sso/internal/domain/auth"
	domain "github.com/isaqueveras/power-sso/internal/domain/auth"
	repo "github.com/isaqueveras/power-sso/internal/infrastructure/auth"
	"github.com/isaqueveras/power-sso/pkg/conversor"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// Register is the business logic for the user register
func Register(ctx context.Context, request *RegisterRequest) error {
	transaction, err := postgres.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	// TODO: find user by email and check if exists or not

	if err = request.Prepare(); err != nil {
		return oops.Err(err)
	}

	var data *auth.Register
	if data, err = conversor.TypeConverter[domain.Register](&request); err != nil {
		return oops.Err(err)
	}

	var repository = repo.New(transaction)
	if err = repository.Register(data); err != nil {
		return oops.Err(err)
	}

	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}
