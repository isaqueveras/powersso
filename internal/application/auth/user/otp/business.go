// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	"context"
	"encoding/base32"

	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/config"
	infraOTP "github.com/isaqueveras/power-sso/internal/infrastructure/auth/user/otp"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/otp"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
	"github.com/isaqueveras/power-sso/pkg/security"
)

// Configure performs business logic to configure otp for a user
func Configure(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *postgres.DBTransaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	data := []byte(security.RandomString(26))
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)

	repository := infraOTP.New(tx)
	if err = repository.Configure(userID, utils.GetStringPointer(string(dst))); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// Unconfigure performs business logic to unconfigure otp for a user
func Unconfigure(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *postgres.DBTransaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repository := infraOTP.New(tx)
	if err = repository.Unconfigure(userID); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// GetQrCode performs business logic to get qrcode url
func GetQrCode(ctx context.Context, userID *uuid.UUID) (res *QRCodeResponse, err error) {
	var (
		tx              *postgres.DBTransaction
		userName, token *string
	)

	if tx, err = postgres.NewTransaction(ctx, true); err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	if userName, token, err = infraOTP.New(tx).GetToken(userID); err != nil {
		return nil, oops.Err(err)
	}

	if config.Get().Server.IsModeDevelopment() {
		*userName += " [DEV]"
	}

	res = new(QRCodeResponse)
	res.Url = utils.GetStringPointer(otp.GetUrlQrCode(*token, *userName))

	return
}
