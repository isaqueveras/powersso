// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package session

import "time"

// Session define a session model
type Session struct {
	ID        *string    `sql:"id"`
	UserID    *string    `sql:"user_id"`
	ExpiresAt *time.Time `sql:"expires_at"`
	CreatedAt *time.Time `sql:"create_at"`
	DeletedAt *time.Time `sql:"deleted_at"`
}
