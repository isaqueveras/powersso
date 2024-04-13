// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	database "github.com/isaqueveras/powersso/database/postgres"
	domain "github.com/isaqueveras/powersso/domain/authentication"
	infra "github.com/isaqueveras/powersso/infrastructure/persistencie/authentication/postgres"
)

var _ domain.IAuth = (*repoAuth)(nil)

type repoAuth struct {
	pg *infra.PGAuth
}

// NewAuthRepository creates a new repository
func NewAuthRepository(tx *database.Transaction) domain.IAuth {
	return &repoAuth{pg: &infra.PGAuth{DB: tx}}
}

// CreateAccount contains the flow for the user register in database
func (r *repoAuth) CreateAccount(data *domain.CreateAccount) (userID *uuid.UUID, err error) {
	return r.pg.CreateAccount(data)
}

// AddAttempts contains the flow for the add number failed attempts
func (r *repoAuth) AddAttempts(userID *uuid.UUID) error {
	return r.pg.AddAttempts(userID)
}

// LoginSteps contains the flow to get the data needed to retrieve the steps required to log in a user
func (r *repoAuth) LoginSteps(email *string) (*domain.Steps, error) {
	return r.pg.LoginSteps(email)
}

var _ domain.IOTP = (*repoOTP)(nil)

type repoOTP struct{ pg *infra.OTP }

// NewOTPRepo creates a new repository
func NewOTPRepo(tx *database.Transaction) domain.IOTP {
	return &repoOTP{pg: &infra.OTP{DB: tx}}
}

func (r *repoOTP) GetToken(userID *uuid.UUID) (*string, *string, error) {
	return r.pg.GetToken(userID)
}

func (r *repoOTP) SetToken(userID *uuid.UUID, secret *string) error {
	return r.pg.SetToken(userID, secret)
}

var _ domain.IFlag = (*repoFlag)(nil)

type repoFlag struct {
	pg *infra.Flag
}

// NewFlagRepo creates a new repository
func NewFlagRepo(tx *database.Transaction) domain.IFlag {
	return &repoFlag{pg: &infra.Flag{DB: tx}}
}

func (r *repoFlag) Set(userID *uuid.UUID, flag domain.Flag) error {
	return r.pg.Set(userID, flag)
}

func (r *repoFlag) Get(userID *uuid.UUID) (*int64, error) {
	return r.pg.Get(userID)
}

var _ domain.ISession = (*repoSession)(nil)

type repoSession struct {
	pg *infra.Session
}

// NewSessionRepository creates a new repository
func NewSessionRepository(tx *database.Transaction) domain.ISession {
	return &repoSession{pg: &infra.Session{DB: tx}}
}

// Create create a new session for a user
func (r *repoSession) Create(userID *uuid.UUID, clientIP, userAgent *string) (*uuid.UUID, error) {
	return r.pg.Create(userID, clientIP, userAgent)
}

// Delete delete a session for a user
func (r *repoSession) Delete(ids ...*uuid.UUID) error {
	return r.pg.Delete(ids...)
}

func (r *repoSession) Get(userID *uuid.UUID) ([]*uuid.UUID, error) {
	return r.pg.Get(userID)
}

var _ domain.IUser = (*repoUser)(nil)

type repoUser struct {
	pg *infra.User
}

// NewUserRepository creates a new repository
func NewUserRepository(tx *database.Transaction) domain.IUser {
	return &repoUser{pg: &infra.User{DB: tx}}
}

// GetUser manages the flow for a user's data
func (r *repoUser) GetUser(user *domain.User) error {
	return r.pg.GetUser(user)
}

// AccountExists manages the flow to check if a user with the same identifier already exists
func (r *repoUser) AccountExists(email *string) error {
	return r.pg.AccountExists(email)
}

// DisableUser manages the flow to deactivate a user's account
func (r *repoUser) DisableUser(userUUID *uuid.UUID) error {
	return r.pg.Disable(userUUID)
}

// ChangePassword manages the flow to change a user's password
func (r *repoUser) ChangePassword(in *domain.ChangePassword) error {
	return r.pg.ChangePassword(in)
}
