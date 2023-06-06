// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/internal/domain/auth"
	"github.com/isaqueveras/powersso/internal/infrastructure/persistencie/auth/mail"
	infra "github.com/isaqueveras/powersso/internal/infrastructure/persistencie/auth/postgres"
	pg "github.com/isaqueveras/powersso/pkg/database/postgres"
	"github.com/isaqueveras/powersso/pkg/mailer"
)

var _ auth.IAuth = (*repoAuth)(nil)

type repoAuth struct {
	pg     *infra.PGAuth
	mailer *mail.MailerAuth
}

// NewAuthRepository creates a new repository
func NewAuthRepository(tx *pg.Transaction, client *mailer.SmtpClient) auth.IAuth {
	return &repoAuth{pg: &infra.PGAuth{DB: tx}, mailer: &mail.MailerAuth{SmtpClient: client, Cfg: config.Get()}}
}

// CreateAccount contains the flow for the user register in database
func (r *repoAuth) CreateAccount(data *auth.CreateAccount) (userID *uuid.UUID, err error) {
	return r.pg.CreateAccount(data)
}

// SendMailActivationAccount contains the flow for the send activation account email
func (r *repoAuth) SendMailActivationAccount(email *string, token *uuid.UUID) error {
	return r.mailer.SendMailActivationAccount(email, token)
}

// GetActivateAccountToken contains the flow for the get activate account token
func (r *repoAuth) GetActivateAccountToken(data *auth.ActivateAccount) error {
	return r.pg.GetActivateAccountToken(data)
}

// CreateAccessToken contains the flow for the create access token
func (r *repoAuth) CreateAccessToken(userID *uuid.UUID) (*uuid.UUID, error) {
	return r.pg.CreateAccessToken(userID)
}

// MarkTokenAsUsed contains the flow for the mark token as used
func (r *repoAuth) MarkTokenAsUsed(token *uuid.UUID) error {
	return r.pg.MarkTokenAsUsed(token)
}

// AddAttempts contains the flow for the add number failed attempts
func (r *repoAuth) AddAttempts(userID *uuid.UUID) error {
	return r.pg.AddAttempts(userID)
}

// LoginSteps contains the flow to get the data needed to retrieve the steps required to log in a user
func (r *repoAuth) LoginSteps(email *string) (*auth.Steps, error) {
	return r.pg.LoginSteps(email)
}
