// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package permission

import "github.com/google/uuid"

// IPermission define an interface for data layer access methods
type IPermission interface {
	Get(userID, organizationID *uuid.UUID) ([]*string, error)
	Create(*Permission) error
}
