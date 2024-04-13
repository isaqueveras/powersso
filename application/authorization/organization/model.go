// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package organization

import "github.com/google/uuid"

// Organization models the data to create a organization
type Organization struct {
	ID          *uuid.UUID `json:"-"`
	Name        *string    `json:"name" binding:"required,min=3,max=20"`
	Description *string    `json:"description"`
	Url         *string    `json:"url" binding:"required,max=200"`
	CreatedByID *uuid.UUID `json:"-"`
}
