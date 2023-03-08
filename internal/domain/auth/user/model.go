// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"time"
)

// Level is the user level
type Level string

const (
	// UserLevel is the user level
	UserLevel Level = "user"
	// AdminLevel is the admin level
	AdminLevel Level = "admin"
)

// User is the model for user
type User struct {
	ID                 *string
	Email              *string
	FirstName          *string
	LastName           *string
	Roles              *string
	About              *string
	UserType           *Level
	Avatar             *string
	PhoneNumber        *string
	Address            *string
	City               *string
	Country            *string
	Gender             *string
	Postcode           *int64
	BlockedTemporarily *bool
	TokenKey           *string
	IsActive           *bool
	OTPToken           *string
	OTPEnabled         *bool
	OTPSetup           *bool
	Birthday           *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
	LoginDate          *time.Time
}

// OTPConfiguredAndEnabled checks if the user has the otp configured and activated for use
func (u *User) OTPConfiguredAndEnabled() bool {
	enabled := u.OTPEnabled != nil && *u.OTPEnabled
	setup := u.OTPSetup != nil && *u.OTPSetup

	if enabled && setup {
		return true
	}

	return false
}
