// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package box

import (
	"time"

	"github.com/google/uuid"
)

// Box models the data of a box
type Box struct {
	ID          *uuid.UUID
	Name        *string
	Desc        *string
	RoomID      *uuid.UUID
	CreatedByID *uuid.UUID
	CreatedAt   *time.Time
}
