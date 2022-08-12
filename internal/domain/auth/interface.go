// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

// IAuth define an interface for data layer access methods
type IAuth interface {
	Register(input *Register) (userID *string, err error)
	SendMailActivationAccount(email *string, token *string) error
	GetActivateAccountToken(token *string) (*ActivateAccountToken, error)
	CreateAccessToken(userID *string) (string, error)
	MarkTokenAsUsed(token *string) error
}
