// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/isaqueveras/power-sso/config"
	appSession "github.com/isaqueveras/power-sso/internal/application/session"
	appUser "github.com/isaqueveras/power-sso/internal/application/user"
	domain "github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	domainSession "github.com/isaqueveras/power-sso/internal/domain/session"
	domainUser "github.com/isaqueveras/power-sso/internal/domain/user"
	"github.com/isaqueveras/power-sso/internal/infrastructure/auth"
	infraRoles "github.com/isaqueveras/power-sso/internal/infrastructure/auth/roles"
	infraSession "github.com/isaqueveras/power-sso/internal/infrastructure/session"
	infraUser "github.com/isaqueveras/power-sso/internal/infrastructure/user"
	"github.com/isaqueveras/power-sso/pkg/conversor"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/mailer"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/pkg/security"
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
		repoUser = infraUser.New(transaction)
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
		repoUser    = infraUser.New(transaction)
		activeToken *domain.ActivateAccountToken
	)

	if activeToken, err = repo.GetActivateAccountToken(token); err != nil {
		return oops.Err(err)
	}

	if *activeToken.Used || !*activeToken.IsValid {
		return oops.Err(ErrTokenIsNotValid())
	}

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

	if err = repo.MarkTokenAsUsed(activeToken.ID); err != nil {
		return oops.Err(err)
	}

	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}

// Login is the business logic for the user login
func Login(ctx context.Context, in *LoginRequest) (*appSession.SessionOut, error) {
	var (
		transaction *postgres.DBTransaction
		err         error
	)

	if transaction, err = postgres.NewTransaction(ctx, false); err != nil {
		return nil, oops.Err(err)
	}
	defer transaction.Rollback()

	var (
		repo        = auth.New(transaction, nil)
		repoUser    = infraUser.New(transaction)
		repoSession = infraSession.New(transaction)
		user        = &domainUser.User{Email: in.Email}
		cfg         = config.Get()

		isAdmin            bool = false
		passw              *string
		timeExpiresSession = time.Now().Add(time.Hour * 24)
		tokenSession       = uuid.NewString()

		// REFACTOR: create function to generate token
		// TODO: add time expires to token in config file
		token, _ = security.NewToken(jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"exp":     timeExpiresSession.Unix(),
		}, cfg.Server.JwtSecretKey, timeExpiresSession.Unix())
	)

	if passw, err = repo.Login(in.Email); err != nil {
		return nil, oops.Err(ErrEmailOrPasswordIsNotValid())
	}

	if err = in.ComparePasswords(passw); err != nil {
		return nil, oops.Err(err)
	}

	if err = repoUser.GetUser(user); err != nil {
		return nil, oops.Err(err)
	}

	if !roles.Exists(roles.CreateSession, roles.Roles{String: *user.Roles}) {
		return nil, oops.Err(ErrNotHavePermissionLogin())
	}

	if err = repoSession.Create(&domainSession.Session{
		UserID:    user.ID,
		Token:     &tokenSession,
		ExpiresAt: &timeExpiresSession,
	}); err != nil {
		return nil, oops.Err(err)
	}

	in.SanitizePassword()
	if err = transaction.Commit(); err != nil {
		return nil, oops.Err(err)
	}

	if roles.Exists(roles.LevelAdmin, roles.Roles{String: *user.Roles}) {
		isAdmin = true
	}

	userRoles := roles.MakeEmptyRoles()
	userRoles.String = *user.Roles
	userRoles.ParseArray()

	return &appSession.SessionOut{
		IsAdmin:   &isAdmin,
		SessionID: &tokenSession,
		User: &appUser.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Roles:     userRoles.Arrays(),
			Avatar:    user.Avatar,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token:     &token,
		ExpiresAt: &timeExpiresSession,
	}, nil
}
