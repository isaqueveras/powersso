// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package roles

// IRoles define an interface for data layer access methods
type IRoles interface {
	RemoveRoles(userID *string, roles string) error
}
