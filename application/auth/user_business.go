// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/google/uuid"
	pg "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/infrastructure/persistencie/auth"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
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

// ChangePassword is the busines logic for change passoword
func ChangePassword(ctx context.Context, in *domain.ChangePassword) (err error) {
	tx, err := pg.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repoUser := auth.NewUserRepository(tx)
	repoSession := auth.NewSessionRepository(tx)

	user := domain.User{ID: in.UserID}
	if err = repoUser.Get(&user); err != nil {
		return oops.Err(err)
	}

	if !user.OTPConfigured() {
		return oops.New("2-factor authentication not configured")
	}

	// Validate code otp
	if err = utils.ValidateToken(user.OTPToken, in.CodeOTP); err != nil {
		return oops.New("2-factor authentication code is invalid")
	}

	// Generate new password crypto
	gen := &domain.CreateAccount{Password: in.Password}
	if err = gen.GeneratePassword(); err != nil {
		return err
	}

	// Change user password
	in.Password, in.Key = gen.Password, gen.Key
	if err = repoUser.ChangePassword(in); err != nil {
		return oops.Err(err)
	}

	// Getting all active user sessions
	var sessions []*uuid.UUID
	if sessions, err = repoSession.Get(in.UserID); err != nil {
		return oops.Err(err)
	}

	// Disabling all active user sessions
	if err = repoSession.Delete(sessions...); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}
