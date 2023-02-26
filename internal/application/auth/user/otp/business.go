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
	transaction, err := postgres.NewTransaction(ctx, false)
	if err != nil {
		return oops.Err(err)
	}
	defer transaction.Rollback()

	data := []byte(security.RandomString(26))
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)

	repository := infraOTP.New(transaction)
	if err = repository.Configure(userID, utils.GetStringPointer(string(dst))); err != nil {
		return oops.Err(err)
	}

	if err = transaction.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// GetQrCode performs business logic to get qrcode url
func GetQrCode(ctx context.Context, userID *uuid.UUID) (*QRCodeResponse, error) {
	transaction, err := postgres.NewTransaction(ctx, true)
	if err != nil {
		return nil, oops.Err(err)
	}
	defer transaction.Rollback()

	var userName, token *string
	if userName, token, err = infraOTP.New(transaction).GetToken(userID); err != nil {
		return nil, oops.Err(err)
	}

	if config.Get().Server.IsModeDevelopment() {
		*userName += " [DEV]"
	}

	url := otp.GetUrlQrCode(*token, *userName)
	return &QRCodeResponse{Url: utils.GetStringPointer(url)}, nil
}
