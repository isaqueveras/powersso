// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mailer

import (
	"fmt"
	"io"
	"net/mail"
	"net/smtp"
	"regexp"
	"strings"

	"github.com/domodwyer/mailyak/v3"
	"github.com/microcosm-cc/bluemonday"

	"github.com/isaqueveras/powersso/config"
)

var _ Mailer = (*SmtpClient)(nil)

// regex to select all tabs
var tabsRegex = regexp.MustCompile(`\t+`)

// SmtpClient defines a SMTP mail client structure that implements
type SmtpClient struct {
	host     string
	port     int
	username string
	password string
	tls      bool
}

// NewSmtpClient creates new `SmtpClient` with the provided configuration.
func NewSmtpClient(host string, port int, username string, password string, tls bool) *SmtpClient {
	return &SmtpClient{
		host:     host,
		port:     port,
		username: username,
		password: password,
		tls:      tls,
	}
}

// Send implements `mailer.Mailer` interface.
func (m *SmtpClient) Send(fromEmail mail.Address, toEmail mail.Address, subject string, htmlBody string, attachments map[string]io.Reader) (err error) {
	var (
		email    *mailyak.MailYak
		smtpAuth = smtp.PlainAuth("", m.username, m.password, m.host)
	)

	if m.tls {
		if email, err = mailyak.NewWithTLS(fmt.Sprintf("%s:%d", m.host, m.port), smtpAuth, nil); err != nil {
			return err
		}
	} else {
		email = mailyak.New(fmt.Sprintf("%s:%d", m.host, m.port), smtpAuth)
	}

	if fromEmail.Name != "" {
		email.FromName(fromEmail.Name)
	}

	email.From(fromEmail.Address)
	email.To(toEmail.Address)
	email.Subject(subject)
	email.HTML().Set(htmlBody)

	policy := bluemonday.StrictPolicy()
	email.Plain().Set(strings.TrimSpace(tabsRegex.ReplaceAllString(policy.Sanitize(htmlBody), "")))

	for name, data := range attachments {
		email.Attach(name, data)
	}

	return email.Send()
}

// Client returns the `SmtpClient` instance.
func Client() *SmtpClient {
	cfg := config.Get()
	return NewSmtpClient(
		cfg.Mailer.Host,
		cfg.Mailer.Port,
		cfg.Mailer.Username,
		cfg.Mailer.Password,
		cfg.Mailer.TLS,
	)
}
