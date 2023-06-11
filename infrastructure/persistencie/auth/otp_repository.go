// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	pg "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/auth"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/auth/postgres"
)

var _ domain.IOTP = (*repoOTP)(nil)

type repoOTP struct {
	pg *infra.PGOTP
}

// NewOTPRepository creates a new repository
func NewOTPRepository(tx *pg.Transaction) domain.IOTP {
	return &repoOTP{pg: &infra.PGOTP{DB: tx}}
}

// GetToken
func (r *repoOTP) GetToken(userID *uuid.UUID) (*string, *string, error) {
	return r.pg.GetToken(userID)
}

// Configure
func (r *repoOTP) Configure(userID *uuid.UUID, secret *string) error {
	return r.pg.Configure(userID, secret)
}

// Unconfigure
func (r *repoOTP) Unconfigure(userID *uuid.UUID) error {
	return r.pg.Unconfigure(userID)
}
