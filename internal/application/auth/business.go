// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"net/http"

	"github.com/isaqueveras/power-sso/config"
	domain "github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	domainUser "github.com/isaqueveras/power-sso/internal/domain/user"
	"github.com/isaqueveras/power-sso/internal/infrastructure/auth"
	infraRoles "github.com/isaqueveras/power-sso/internal/infrastructure/auth/roles"
	"github.com/isaqueveras/power-sso/internal/infrastructure/user"
	"github.com/isaqueveras/power-sso/pkg/conversor"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/mailer"
	"github.com/isaqueveras/power-sso/pkg/oops"
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
		userID   *string
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
	if userID, err = repo.Register(data); err != nil {
		return oops.Err(err)
	}

	var accessToken string
	if accessToken, err = repo.CreateAccessToken(userID); err != nil {
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

// Activation is the business logic for the user activation
func Activation(ctx context.Context, token *string) (err error) {
	transaction, err := postgres.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	var (
		repo        = auth.New(transaction, nil)
		repoUser    = user.New(transaction)
		activeToken *domain.ActivateAccountToken
	)

	if activeToken, err = repo.GetActivateAccountToken(token); err != nil {
		return oops.Err(err)
	}

	if *activeToken.Used || !*activeToken.IsValid {
		return oops.Err(&oops.Error{
			Message:    "Token is not valid",
			StatusCode: http.StatusBadRequest,
		})
	}

	if !*activeToken.Used && *activeToken.IsValid {
		var user = domainUser.User{
			ID: activeToken.UserID,
		}

		if err = repoUser.GetUser(&user); err != nil {
			return oops.Err(err)
		}

		if !roles.Exists(roles.ReadActivationToken, roles.Roles{String: *user.Roles}) {
			return oops.Err(ErrNotHavePermissionActiveAccount())
		}

		var repoRoles = infraRoles.New(transaction)
		if err = repoRoles.RemoveRoles(user.ID, roles.ReadActivationToken); err != nil {
			return oops.Err(err)
		}

		rolesSession := roles.MakeEmptyRoles()
		rolesSession.Add(roles.ReadSession, roles.CreateSession)
		rolesSession.ParseString()

		if err = repoRoles.AddRoles(user.ID, rolesSession.Strings()); err != nil {
			return oops.Err(err)
		}

		// TODO: mark the token as used
	}

	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}
