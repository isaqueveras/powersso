// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	"context"

	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/config"
	infraOTP "github.com/isaqueveras/power-sso/internal/infrastructure/auth/user/otp"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/otp"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// Configure performs business logic to configure otp for a user
func Configure(_ context.Context) (err error) {
	return
}

// GetQRCode performs business logic to get qrcode url
func GetQRCode(ctx context.Context, userID *uuid.UUID) (*QRCodeResponse, error) {
	tx, err := postgres.NewTransaction(ctx, true)
	if err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	var userName, token *string
	if userName, token, err = infraOTP.New(tx).GetToken(userID); err != nil {
		return nil, oops.Err(err)
	}

	if config.Get().Server.IsModeDevelopment() {
		*userName += " [DEV]"
	}

	url := otp.GetUrlQrCode(*token, *userName)
	return &QRCodeResponse{Url: utils.GetStringPointer(url)}, nil
}
