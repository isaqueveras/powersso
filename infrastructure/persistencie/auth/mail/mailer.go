// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mail

import (
	"net/mail"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/mailer"
)

// MailerAuth is the implementation
// of mailer for the auth repository
type MailerAuth struct {
	SmtpClient *mailer.SmtpClient
	Cfg        *config.Config
}

// SendMailActivationAccount send the activation account email
func (ma *MailerAuth) SendMailActivationAccount(email *string, token *uuid.UUID) error {
	return ma.SmtpClient.Send(
		mail.Address{Name: ma.Cfg.Mailer.Username, Address: ma.Cfg.Mailer.Email},
		mail.Address{Address: *email},
		"Activate your "+ma.Cfg.Meta.ProjectName+" registration",
		`Click on the link below to activate your `+ma.Cfg.Meta.ProjectName+` registration:

		<a href="`+ma.Cfg.Meta.ProjectURL+`/auth/activation/`+token.String()+`">`+ma.Cfg.Meta.ProjectURL+`/auth/activation/`+token.String()+`</a>

		If you have not made this request, please ignore this email.

		Yours sincerely,
		`+ma.Cfg.Meta.ProjectName+` team`, nil,
	)
}
