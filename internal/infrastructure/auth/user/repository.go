// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/power-sso/internal/domain/auth/user"
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

// GetUser contains the flow for the get user
func (r *repository) GetUser(data *user.User) error {
	return r.pg.getUser(data)
}

// DisableUser contains the flow for disable user
func (r *repository) DisableUser(userUUID *uuid.UUID) error {
	return r.pg.disableUser(userUUID)
}
