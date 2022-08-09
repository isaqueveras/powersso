// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/isaqueveras/power-sso/config"
	domain "github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	"github.com/isaqueveras/power-sso/internal/infrastructure/auth"
	"github.com/isaqueveras/power-sso/internal/infrastructure/user"
	"github.com/isaqueveras/power-sso/pkg/conversor"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/mailer"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/tokens"
)

// Register is the business logic for the user register
func Register(ctx context.Context, in *RegisterRequest) error {
	cfg := config.Get()

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

	var (
		exists   bool
		data     *domain.Register
		repo     = auth.New(transaction, mailer.Client(cfg))
		repoUser = user.New(transaction)
	)

	if exists, err = repoUser.FindByEmailUserExists(in.Email); err != nil {
		return oops.Err(err)
	}

	if exists {
		return oops.Err(ErrUserExists())
	}

	if data, err = conversor.TypeConverter[domain.Register](&in); err != nil {
		return oops.Err(err)
	}

	data.Roles = &in.Roles.String
	if err = repo.Register(data); err != nil {
		return oops.Err(err)
	}

	var accessToken string
	if accessToken, err = tokens.NewUserVerifyToken(cfg, in.Email, in.TokenKey); err != nil {
		return oops.Err(err)
	}

	if err = repo.SendMailActivationAccount(in.Email, &accessToken); err != nil {
		return oops.Err(err)
	}

	in.SanitizePassword()
	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}
