// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import "github.com/google/uuid"

// IUser define an interface for data layer access methods
type IUser interface {
	GetUser(data *User) error
	FindByEmailUserExists(email *string) (exists bool, err error)
	DisableUser(userUUID *uuid.UUID) error
}
