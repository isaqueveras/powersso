// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"fmt"

	"github.com/isaqueveras/power-sso/config"
	domain "github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	"github.com/isaqueveras/power-sso/internal/infrastructure/auth"
	"github.com/isaqueveras/power-sso/internal/infrastructure/user"
	"github.com/isaqueveras/power-sso/pkg/conversor"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/tokens"
)

// Register is the business logic for the user register
func Register(ctx context.Context, in *RegisterRequest) error {
	transaction, err := postgres.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	if err = in.Prepare(); err != nil {
		return oops.Err(err)
	}

	if in.Roles == nil {
		in.Roles = new(roles.Roles)
		in.Roles.Add(roles.LevelUser, roles.ReadActivationToken)
	}

	in.Roles.Parse()

	var exists bool
	if exists, err = user.
		New(transaction).
		FindByEmailUserExists(in.Email); err != nil {
		return oops.Err(err)
	}

	if exists {
		return oops.Err(ErrUserExists())
	}

	var data *domain.Register
	if data, err = conversor.TypeConverter[domain.Register](&in); err != nil {
		return oops.Err(err)
	}

	data.Roles = &in.Roles.String
	if err = auth.
		New(transaction).
		Register(data); err != nil {
		return oops.Err(err)
	}

	var accessToken string
	if accessToken, err = tokens.NewUserVerifyToken(config.Get(), in.Email, in.TokenKey); err != nil {
		return oops.Err(err)
	}

	fmt.Printf("accessToken: %v\n", accessToken)

	// TODO: send email to user with the confirmation link

	in.SanitizePassword()
	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}
