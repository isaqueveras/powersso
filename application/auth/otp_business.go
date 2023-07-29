// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"encoding/base32"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/infrastructure/persistencie/auth"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// Configure performs business logic to configure otp for a user
func Configure(ctx context.Context, userID *uuid.UUID) (err error) {
	var tx *postgres.Transaction
	if tx, err = postgres.NewTransaction(ctx, false); err != nil {
		return oops.Err(err)
	}
	defer tx.Rollback()

	data := []byte(utils.RandomString(26))
	dst := make([]byte, base32.StdEncoding.EncodedLen(len(data)))
	base32.StdEncoding.Encode(dst, data)

	repo := auth.NewOTPRepository(tx)
	if err = repo.Configure(userID, utils.Pointer(string(dst))); err != nil {
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

	res = &domain.QRCode{Url: utils.Pointer(utils.GetUrlQrCode(*token, *userName))}
	return
}
