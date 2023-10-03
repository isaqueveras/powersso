// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/infrastructure/persistencie/auth"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// Configure2FA performs business logic to configure otp for a user
func Configure2FA(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	service := domain.NewAuthService(auth.NewFlagRepo(tx), auth.NewOTPRepo(tx, userID))
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
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repoFlag := auth.NewFlagRepo(tx)
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

	repoOTP := auth.NewOTPRepo(tx, userID)
	if err = repoOTP.SetToken(nil); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// GetQRCode2FA performs business logic to get qrcode url
func GetQRCode2FA(ctx context.Context, userID *uuid.UUID) (url *string, err error) {
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, true); err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	var userName, token *string
	if userName, token, err = auth.NewOTPRepo(tx, userID).GetToken(); err != nil {
		return nil, oops.Err(err)
	}

	if config.Get().Server.IsModeDevelopment() {
		*userName += " [DEV]"
	}

	return utils.Pointer(utils.GetUrlQrCode(*token, *userName)), nil
}
