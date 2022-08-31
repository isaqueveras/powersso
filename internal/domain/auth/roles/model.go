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

	CreateSession string = "create:session"
	ReadSession   string = "read:session"
	DeleteSession string = "delete:session"
)

// Roles type for user roles
type Roles struct {
	Array  []string
	String string
}

// MakeEmptyRoles make an empty roles slice
func MakeEmptyRoles() *Roles {
	return &Roles{}
}

// Strings return the roles slice as a string
func (r *Roles) Strings() string {
	return r.String
}

// Arrays return the roles slice as list of strings
func (r *Roles) Arrays() []string {
	return r.Array
}

// Parse parse the roles string to a slice of strings
func (r *Roles) Parse() {
	r.ParseString()
	r.ParseArray()
}

// Exists check if the role exists in the roles slice
func Exists(value string, roles Roles) bool {
	if roles.Array != nil {
		for _, r := range roles.Array {
			if r == value {
				return true
			}
		}
	}

	if roles.String != "" {
		if strings.Contains(roles.String, value) {
			return true
		}
	}

	return false
}

// Remove remove the role from the roles slice
func (r *Roles) Remove(role string) {
	if !Exists(role, Roles{Array: r.Array}) {
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
	r.ParseString()
}

// Add add the role to the roles slice
func (r *Roles) Add(role ...string) {
	r.Array = append(r.Array, role...)
	r.ParseString()
}

// ParseString parse the roles slice to a string
func (r *Roles) ParseString() {
	if r.Array != nil {
		var _temp string
		for i := range r.Array {
			_temp += r.Array[i] + ","
		}
		r.String = fmt.Sprintf("{%s}", strings.Trim(_temp, ","))
	}
}

// ParseArray parse the roles string to a slice of strings
func (r *Roles) ParseArray() {
	if r.String != "" {
		r.Array = strings.Split(strings.TrimSuffix(strings.TrimPrefix(r.String, "{"), "}"), ",")
	}
}
