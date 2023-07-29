// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import "github.com/google/uuid"

// IAuth define an interface for data layer access methods
type IAuth interface {
	CreateAccount(*CreateAccount) (userID *uuid.UUID, err error)
	SendMailActivationAccount(email *string, token *uuid.UUID) error
	GetActivateAccountToken(data *ActivateAccount) error
	CreateAccessToken(userID *uuid.UUID) (*uuid.UUID, error)
	MarkTokenAsUsed(token *uuid.UUID) error
	AddAttempts(userID *uuid.UUID) error
	LoginSteps(email *string) (*Steps, error)
}

// ISession define an interface for data layer access methods
type ISession interface {
	Create(userID *uuid.UUID, clientIP, userAgent *string) (*uuid.UUID, error)
	Delete(sessionID *uuid.UUID) error
}

// IFlag define an interface for data layer access methods
type IFlag interface {
	Get(userID *uuid.UUID) (*int64, error)
	Set(userID *uuid.UUID, flag *Flag) error
}

// IOTP define an interface for data layer access methods
type IOTP interface {
	GetToken() (*string, *string, error)
	SetToken(secret *string) error
}

// IUser define an interface for data layer access methods
type IUser interface {
	Get(user *User) error
	Exist(email *string) error
	Disable(userUUID *uuid.UUID) error
}
