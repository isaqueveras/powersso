// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"time"
)

// User is the model for user
type User struct {
	ID          *string
	Email       *string
	FirstName   *string
	LastName    *string
	Roles       *string
	About       *string
	Avatar      *string
	PhoneNumber *string
	Address     *string
	City        *string
	Country     *string
	Gender      *string
	Postcode    *int
	TokenKey    *string
	Birthday    *time.Time
	CreateAt    *time.Time
	UpdateAt    *time.Time
	LoginDate   *time.Time
}
