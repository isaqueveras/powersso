// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"encoding/base32"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	domain "github.com/isaqueveras/powersso/internal/domain/auth"
	"github.com/isaqueveras/powersso/internal/infrastructure/persistencie/auth"
	"github.com/isaqueveras/powersso/internal/utils"
	"github.com/isaqueveras/powersso/otp"
	"github.com/isaqueveras/powersso/pkg/database/postgres"
	"github.com/isaqueveras/powersso/pkg/oops"
	"github.com/isaqueveras/powersso/pkg/security"
)

// Configure performs business logic to configure otp for a user
func Configure(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	data := []byte(security.RandomString(26))
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)

	repo := auth.NewOTPRepository(tx)
	if err = repo.Configure(userID, utils.GetStringPointer(string(dst))); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// Unconfigure performs business logic to unconfigure otp for a user
func Unconfigure(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	repository := auth.NewOTPRepository(tx)
	if err = repository.Unconfigure(userID); err != nil {
		return oops.Err(err)
	}

	if err = tx.Commit(); err != nil {
		return oops.Err(err)
	}

	return
}

// GetQrCode performs business logic to get qrcode url
func GetQrCode(ctx context.Context, userID *uuid.UUID) (res *domain.QRCode, err error) {
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, true); err != nil {
		return nil, oops.Err(err)
	}
	defer tx.Rollback()

	var userName, token *string
	if userName, token, err = auth.NewOTPRepository(tx).GetToken(userID); err != nil {
		return nil, oops.Err(err)
	}

	if config.Get().Server.IsModeDevelopment() {
		*userName += " [DEV]"
	}

	res = &domain.QRCode{Url: utils.GetStringPointer(otp.GetUrlQrCode(*token, *userName))}
	return
}
