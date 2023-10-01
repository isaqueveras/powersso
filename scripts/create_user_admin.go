// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package scripts

import (
	"context"
	"log"
	"time"

	"github.com/isaqueveras/powersso/application/auth"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/utils"
)

// CreateUserAdmin register the first admin user
func CreateUserAdmin(logg *utils.Logger) {
	logg.Info("Initializing script create user admin")
	defer logg.Info("Finalizing script create user admin")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	tx, err := postgres.NewTransaction(ctx, false)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	if err = auth.CreateAccount(ctx, &domain.CreateAccount{
		FirstName: utils.Pointer("User Power"),
		LastName:  utils.Pointer("Admin"),
		Email:     utils.Pointer("admin@powersso.io"),
		Password:  utils.Pointer("admin123456"),
		Level:     utils.Pointer(domain.AdminLevel),
	}); err != nil {
		if err.Error() == domain.ErrUserExists().Error() {
			return
		}
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
