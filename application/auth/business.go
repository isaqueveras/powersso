// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/google/uuid"
	database "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/auth"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/tokens"
	"github.com/isaqueveras/powersso/utils"
)

// CreateAccount is the business logic for the user register
func CreateAccount(ctx context.Context, in *domain.CreateAccount) (url *string, err error) {
	var tx *database.Transaction
	if tx, err = database.NewTransaction(ctx, false); err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	if err = in.Prepare(); err != nil {
		return nil, oops.Err(err)
	}

	userRepository := infra.NewUserRepository(tx)
	if err = userRepository.AccountExists(in.Email); err != nil {
		return nil, oops.Err(err)
	}

	authRepository := infra.NewAuthRepository(tx)
	var userID *uuid.UUID
	if userID, err = authRepository.CreateAccount(in); err != nil {
		return nil, oops.Err(err)
	}

	service := domain.NewAuthService(infra.NewFlagRepo(tx), infra.NewOTPRepo(tx))
	if err = service.Configure2FA(userID); err != nil {
		return nil, oops.Err(err)
	}

	if url, err = service.GenerateQrCode2FA(userID); err != nil {
		return nil, oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, oops.Err(err)
	}

	return
}

// Login is the business logic for the user login
func Login(ctx context.Context, in *domain.Login) (*domain.Session, error) {
	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	var (
		repoAuth    = infra.NewAuthRepository(tx)
		repoUser    = infra.NewUserRepository(tx)
		repoSession = infra.NewSessionRepository(tx)
	)

	user := &domain.User{Email: in.Email}
	if err = repoUser.GetUser(user); err != nil {
		return nil, oops.Err(err)
	}

	if !user.IsActive() {
		return nil, domain.ErrUserNotExists()
	}

	if !user.OTPConfigured() {
		return nil, domain.ErrAuthentication2factorNotConfigured()
	}

	if user.IsBlocked() {
		return nil, domain.ErrUserBlockedTemporarily()
	}

	if err = in.ComparePasswords(user.Password, user.Key); err != nil {
		if errAttempts := repoAuth.AddAttempts(user.ID); errAttempts != nil {
			return nil, oops.Err(errAttempts)
		}
		if errAttempts := tx.Commit(); errAttempts != nil {
			return nil, oops.Err(errAttempts)
		}
		return nil, oops.Err(err)
	}

	if err = utils.ValidateToken(user.OTPToken, in.OTP); err != nil {
		return nil, domain.ErrOTPTokenInvalid()
	}

	var sessionID *uuid.UUID
	if sessionID, err = repoSession.Create(user.ID, in.ClientIP, in.UserAgent); err != nil {
		return nil, oops.Err(err)
	}

	var token *string
	if token, err = tokens.NewUserAuthToken(user, sessionID); err != nil {
		return nil, oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return nil, oops.Err(err)
	}

	return &domain.Session{
		SessionID: sessionID,
		Level:     user.Level,
		UserID:    user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		Token:     token,
	}, nil
}

// Logout is the business logic for the user logout
func Logout(ctx context.Context, sessionID *uuid.UUID) (err error) {
	var tx *database.Transaction
	if tx, err = database.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	if err = infra.NewSessionRepository(tx).Delete(sessionID); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// LoginSteps is the business logic needed to retrieve needed steps for log a user in
func LoginSteps(ctx context.Context, email *string) (res *domain.Steps, err error) {
	var tx *database.Transaction
	if tx, err = database.NewTransaction(ctx, true); err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	return infra.NewAuthRepository(tx).LoginSteps(email)
}

// Configure2FA performs business logic to configure otp for a user
func Configure2FA(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *database.Transaction
	if tx, err = database.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	var (
		repoOTP  = infra.NewOTPRepo(tx)
		repoFlag = infra.NewFlagRepo(tx)
		service  = domain.NewAuthService(repoFlag, repoOTP)
	)

	if err = service.Configure2FA(userID); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// Unconfigure2FA performs business logic to unconfigure otp for a user
func Unconfigure2FA(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *database.Transaction
	if tx, err = database.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repoFlag := infra.NewFlagRepo(tx)
	flag, err := repoFlag.Get(userID)
	if err != nil {
		return oops.Err(err)
	}

	if err = repoFlag.Set(userID, (domain.Flag(*flag))&(^domain.FlagOTPEnable)); err != nil {
		return oops.Err(err)
	}

	if err = repoFlag.Set(userID, (domain.Flag(*flag))&(^domain.FlagOTPSetup)); err != nil {
		return oops.Err(err)
	}

	repoOTP := infra.NewOTPRepo(tx)
	if err = repoOTP.SetToken(userID, nil); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// GetQRCode2FA performs business logic to get qrcode url
func GetQRCode2FA(ctx context.Context, userID *uuid.UUID) (url *string, err error) {
	var tx *database.Transaction
	if tx, err = database.NewTransaction(ctx, true); err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	var (
		repoFlag = infra.NewFlagRepo(tx)
		repoOTP  = infra.NewOTPRepo(tx)
		service  = domain.NewAuthService(repoFlag, repoOTP)
	)

	return service.GenerateQrCode2FA(userID)
}

// DisableUser is the business logic for disable user
func DisableUser(ctx context.Context, userUUID *uuid.UUID) error {
	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repo := infra.NewUserRepository(tx)
	if err = repo.DisableUser(userUUID); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return nil
}

// ChangePassword is the busines logic for change passoword
func ChangePassword(ctx context.Context, in *domain.ChangePassword) (err error) {
	tx, err := database.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repoUser := infra.NewUserRepository(tx)
	repoSession := infra.NewSessionRepository(tx)

	user := domain.User{ID: in.UserID}
	if err = repoUser.GetUser(&user); err != nil {
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
