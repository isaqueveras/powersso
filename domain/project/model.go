// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

// CreateProject models the data to create a project
type CreateProject struct {
	Name        *string
	Description *string
	Slug        *string
	Url         *string
	CreatedByID *string
}
