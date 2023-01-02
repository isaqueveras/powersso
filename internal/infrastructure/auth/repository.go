// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/mailer"
)

var _ auth.IAuth = (*repository)(nil)

// repository is the implementation of the auth repository
type repository struct {
	pg     *pgAuth
	mailer *mailerAuth
}

// New creates a new repository
func New(transaction *postgres.DBTransaction, smtpClient *mailer.SmtpClient) auth.IAuth {
	cfg := config.Get()

	return &repository{
		pg: &pgAuth{
			DB: transaction,
		},
		mailer: &mailerAuth{
			smtpClient: smtpClient,
			cfg:        cfg,
		},
	}
}

// Register contains the flow for the user register in database
func (r *repository) Register(input *auth.Register) (userID *string, err error) {
	return r.pg.register(input)
}

// SendMailActivationAccount contains the flow for the send activation account email
func (r *repository) SendMailActivationAccount(email *string, token *string) error {
	if config.Get().Server.Mode == config.ModeTesting {
		return nil
	}
	return r.mailer.sendMailActivationAccount(email, token)
}

// GetActivateAccountToken contains the flow for the get activate account token
func (r *repository) GetActivateAccountToken(token *string) (*auth.ActivateAccountToken, error) {
	return r.pg.getActivateAccountToken(token)
}

// CreateAccessToken contains the flow for the create access token
func (r *repository) CreateAccessToken(userID *string) (string, error) {
	return r.pg.createAccessToken(userID)
}

// MarkTokenAsUsed contains the flow for the mark token as used
func (r *repository) MarkTokenAsUsed(token *string) error {
	return r.pg.markTokenAsUsed(token)
}

// Login contains the flow for the user login
func (r *repository) Login(email *string) (*string, error) {
	return r.pg.login(email)
}

// AddNumberFailedAttempts contains the flow for the add number failed attempts
func (r *repository) AddNumberFailedAttempts(userID *string) error {
	return r.pg.addNumberFailedAttempts(userID)
}
