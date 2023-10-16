// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import "github.com/google/uuid"

// IAuthService defines an interface for service methods to access the data layer
type IAuthService interface {
	Configure2FA(userID *uuid.UUID) error
	GenerateQrCode2FA(userID *uuid.UUID) (*string, error)
}

// IAuth define an interface for data layer access methods
type IAuth interface {
	CreateAccount(*CreateAccount) (userID *uuid.UUID, err error)
	AddAttempts(userID *uuid.UUID) error
	LoginSteps(email *string) (*Steps, error)
}

// ISession define an interface for data layer access methods
type ISession interface {
	Create(userID *uuid.UUID, clientIP, userAgent *string) (*uuid.UUID, error)
	Delete(ids ...*uuid.UUID) error
	Get(userID *uuid.UUID) ([]*uuid.UUID, error)
}

// IFlag define an interface for data layer access methods
type IFlag interface {
	Get(userID *uuid.UUID) (*int64, error)
	Set(userID *uuid.UUID, flag Flag) error
}

// IOTP define an interface for data layer access methods
type IOTP interface {
	GetToken(userID *uuid.UUID) (*string, *string, error)
	SetToken(userID *uuid.UUID, secret *string) error
}

// IUser define an interface for data layer access methods
type IUser interface {
	GetUser(*User) error
	ChangePassword(*ChangePassword) error
	AccountExists(email *string) error
	DisableUser(userUUID *uuid.UUID) error
}
