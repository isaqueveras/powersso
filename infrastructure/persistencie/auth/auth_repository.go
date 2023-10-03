// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	pg "github.com/isaqueveras/powersso/database/postgres"
	"github.com/isaqueveras/powersso/domain/auth"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/auth/postgres"
)

var _ auth.IAuth = (*repoAuth)(nil)

type repoAuth struct{ pg *infra.PGAuth }

// NewAuthRepository creates a new repository
func NewAuthRepository(tx *pg.Transaction) auth.IAuth {
	return &repoAuth{pg: &infra.PGAuth{DB: tx}}
}

// CreateAccount contains the flow for the user register in database
func (r *repoAuth) CreateAccount(data *auth.CreateAccount) (userID *uuid.UUID, err error) {
	return r.pg.CreateAccount(data)
}

// AddAttempts contains the flow for the add number failed attempts
func (r *repoAuth) AddAttempts(userID *uuid.UUID) error {
	return r.pg.AddAttempts(userID)
}

// LoginSteps contains the flow to get the data needed to retrieve the steps required to log in a user
func (r *repoAuth) LoginSteps(email *string) (*auth.Steps, error) {
	return r.pg.LoginSteps(email)
}
