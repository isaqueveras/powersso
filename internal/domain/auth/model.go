// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"time"
)

// Register model the data to register user in the database
type Register struct {
	FirstName   *string    `sql:"first_name" json:"first_name"`
	LastName    *string    `sql:"last_name" json:"last_name"`
	Email       *string    `sql:"email" json:"email"`
	Password    *string    `sql:"password" json:"password"`
	Roles       *string    `sql:"roles" json:"-"`
	About       *string    `sql:"about" json:"about"`
	Avatar      *string    `sql:"avatar" json:"avatar"`
	PhoneNumber *string    `sql:"phone_number" json:"phone_number"`
	Address     *string    `sql:"address" json:"address"`
	City        *string    `sql:"city" json:"city"`
	Country     *string    `sql:"country" json:"country"`
	Gender      *string    `sql:"gender" json:"gender"`
	Postcode    *int       `sql:"postcode" json:"postcode"`
	Birthday    *time.Time `sql:"birthday" json:"birthday"`
	TokenKey    *string    `sql:"token_key" json:"token_key"`
}

// ActivateAccountToken model the data to activate user account
type ActivateAccountToken struct {
	ID        *string
	UserID    *string
	Used      *bool
	IsValid   *bool
	ExpiresAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
