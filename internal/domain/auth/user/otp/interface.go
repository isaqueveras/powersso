// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import "github.com/google/uuid"

// IOTP define an interface for data layer access methods
type IOTP interface {
	GetToken(userID *uuid.UUID) (*string, *string, error)
	Configure(userID *uuid.UUID, secret *string) error
	Unconfigure(userID *uuid.UUID) error
}
