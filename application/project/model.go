// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

// NewProject models the data to create a project
type NewProject struct {
	Name        *string `json:"name" binding:"required,min=3,max=20"`
	Description *string `json:"description"`
	Slug        *string `json:"-"`
	Url         *string `json:"url" binding:"required,max=200"`
	CreatedByID *string `json:"-"`
}
