// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/isaqueveras/power-sso/internal/domain/user"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
)

var _ user.IUser = (*repository)(nil)

// repository is the implementation of the user repository
type repository struct {
	pg *pgUser
}

// New creates a new repository
func New(transaction *postgres.DBTransaction) user.IUser {
	return &repository{pg: &pgUser{DB: transaction}}
}

// FindByEmailUserExists contains the flow for the find user by email in database
func (r *repository) FindByEmailUserExists(email *string) (bool, error) {
	return r.pg.findByEmailUserExists(email)
}
