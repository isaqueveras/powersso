// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package mailer

import (
	"io"
	"net/mail"
)

// Mailer defines a base mail client interface.
type Mailer interface {
	// Send sends an email with HTML body to the specified recipient.
	Send(fromEmail mail.Address, toEmail mail.Address, subject string, htmlBody string, attachments map[string]io.Reader) error
}
