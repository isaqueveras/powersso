// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package organization

import (
	"time"

	"github.com/google/uuid"
)

// Organization models the data of a organization
type Organization struct {
	ID          *uuid.UUID
	Name        *string
	Desc        *string
	URL         *string
	CreatedByID *uuid.UUID
	CreatedAt   *time.Time
}
