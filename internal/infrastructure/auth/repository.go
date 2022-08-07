// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
)

type repository struct {
	pg *pgAuth
}

// New creates a new repository
func New(transaction *postgres.DBTransaction) auth.IAuth {
	return &repository{pg: &pgAuth{DB: transaction}}
}

// Register contains the flow for the user register in database
func (r *repository) Register(input *auth.Register) error {
	return r.pg.register(input)
}
