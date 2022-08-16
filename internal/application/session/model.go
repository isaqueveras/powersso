// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package session

import (
	"time"

	"github.com/isaqueveras/power-sso/internal/application/user"
)

// SessionOut define a session model
// output for presentation layer
type SessionOut struct {
	SessionID *string    `json:"session_id,omitempty"`
	IsAdmin   *bool      `json:"is_admin,omitempty"`
	User      *user.User `json:"user,omitempty"`
	Token     *string    `json:"token_jwt,omitempty"`
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}
