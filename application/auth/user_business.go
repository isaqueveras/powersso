// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/infrastructure/persistencie/auth"
	"github.com/isaqueveras/powersso/oops"
)

// Disable is the business logic for disable user
func Disable(ctx context.Context, userUUID *uuid.UUID) error {
	tx, err := pg.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repo := auth.NewUserRepository(tx)
	if err = repo.Disable(userUUID); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}
