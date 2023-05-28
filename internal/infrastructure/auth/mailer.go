// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package auth

import (
	"net/mail"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/pkg/mailer"
)

// mailerAuth is the implementation
// of mailer for the auth repository
type mailerAuth struct {
	smtpClient *mailer.SmtpClient
	cfg        *config.Config
}

// SendMailActivationAccount send the activation account email
func (ma *mailerAuth) sendMailActivationAccount(email *string, token *string) error {
	return ma.smtpClient.Send(
		mail.Address{Name: ma.cfg.Mailer.Username, Address: ma.cfg.Mailer.Email},
		mail.Address{Address: *email},
		"Activate your "+ma.cfg.Meta.ProjectName+" registration",
		`Click on the link below to activate your `+ma.cfg.Meta.ProjectName+` registration:

		<a href="`+ma.cfg.Meta.ProjectURL+`/auth/activation/`+*token+`">`+ma.cfg.Meta.ProjectURL+`/auth/activation/`+*token+`</a>

		If you have not made this request, please ignore this email.

		Yours sincerely,
		`+ma.cfg.Meta.ProjectName+` team`,
		nil,
	)
}
