// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

// IProject define an interface for data layer access methods
type IProject interface {
	CreateNewProject(*CreateProject) error
}
