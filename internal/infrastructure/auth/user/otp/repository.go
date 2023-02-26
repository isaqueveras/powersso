// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/internal/domain/auth/user/otp"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
)

// repository is the implementation of the otp repository
type repository struct {
	pg *PGOTP
}

// New creates a new repository
func New(transaction *postgres.DBTransaction) otp.IOTP {
	return &repository{pg: &PGOTP{DB: transaction}}
}

// GetToken return the token of a user's otp
func (r *repository) GetToken(userID *uuid.UUID) (*string, *string, error) {
	return r.pg.GetToken(userID)
}

// Configure configure otp for a user
func (r *repository) Configure(userID *uuid.UUID, secret *string) error {
	return r.pg.Configure(userID, secret)
}
