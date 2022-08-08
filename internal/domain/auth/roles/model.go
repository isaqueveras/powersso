// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package roles

import (
	"fmt"
	"strings"
)

// Constants for the roles
const (
	// UserRole is the user role
	LevelUser string = "level:user"
	// LevelAdmin is the admin role
	LevelAdmin string = "level:admin"

	// ReadActivationToken is the read activation token role
	ReadActivationToken string = "read:activation_token"
)

// Roles type for user roles
type Roles struct {
	Array  []string
	String string
}

// Parse parse the roles string to a slice of strings
func (r *Roles) Parse() {
	r.parseString()

	if r.Array == nil {
		r.Array = strings.Split(r.String, ",")
	}
}

// Exists check if the role exists in the roles slice
func (r *Roles) Exists(role string) bool {
	for _, r := range r.Array {
		if r == role {
			return true
		}
	}
	return false
}

// Remove remove the role from the roles slice
func (r *Roles) Remove(role string) {
	if !r.Exists(role) {
		return
	}

	var _temp []string
	for i := range r.Array {
		if r.Array[i] == role {
			r.Array[i] = ""
		}

		if r.Array[i] != "" {
			_temp = append(_temp, r.Array[i])
		}
	}

	r.Array = nil
	r.Array = _temp
	r.parseString()
}

// Add add the role to the roles slice
func (r *Roles) Add(role ...string) {
	r.Array = append(r.Array, role...)
	r.parseString()
}

func (r *Roles) parseString() {
	var _temp string
	for i := range r.Array {
		_temp += r.Array[i] + ","
	}
	r.String = fmt.Sprintf("{%s}", strings.Trim(_temp, ","))
}
