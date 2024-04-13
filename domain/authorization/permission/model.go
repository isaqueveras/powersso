// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package permission

import "github.com/google/uuid"

// Permission ...
type Permission struct {
	Name        *string
	Credential  *string
	CreatedByID *uuid.UUID
}
