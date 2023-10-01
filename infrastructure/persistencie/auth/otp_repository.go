// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/auth/postgres"
)

var _ domain.IOTP = (*repoOTP)(nil)

type repoOTP struct{ pg *infra.PGOTP }

func NewOTPRepo(tx *postgres.Transaction, userID *uuid.UUID) domain.IOTP {
	return &repoOTP{pg: &infra.PGOTP{DB: tx, UserID: userID}}
}

func (r *repoOTP) GetToken() (*string, *string, error) {
	return r.pg.GetToken()
}

func (r *repoOTP) SetToken(secret *string) error {
	return r.pg.SetToken(secret)
}