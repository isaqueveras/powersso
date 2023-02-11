// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/isaqueveras/power-sso/internal/infrastructure/auth/user"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// Disable is the business logic for disable user
func Disable(ctx context.Context, userUUID *uuid.UUID) (err error) {
	var transaction *postgres.DBTransaction
	if transaction, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	repository := user.New(transaction)
	if err = repository.DisableUser(userUUID); err != nil {
		return oops.Err(err)
	}

	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}
