// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	domain "github.com/isaqueveras/powersso/internal/domain/auth"
	infra "github.com/isaqueveras/powersso/internal/infrastructure/persistencie/auth/postgres"
	pg "github.com/isaqueveras/powersso/pkg/database/postgres"
)

var _ domain.IUser = (*repoUser)(nil)

type repoUser struct{ pg *infra.PGUser }

// NewUserRepository creates a new repository
func NewUserRepository(tx *pg.Transaction) domain.IUser {
	return &repoUser{pg: &infra.PGUser{DB: tx}}
}

// Get get user data
func (r *repoUser) Get(user *domain.User) error {
	return r.pg.Get(user)
}

// Exist check if user already exists
func (r *repoUser) Exist(email *string) error {
	return r.pg.Exist(email)
}

// Disable deactivate a user's account
func (r *repoUser) Disable(userUUID *uuid.UUID) error {
	return r.pg.Disable(userUUID)
}
