// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"

	"github.com/isaqueveras/power-sso/pkg/i18n"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// ErrUserExists creates and returns an error when the user already exists
func ErrUserExists() *oops.Error {
	return oops.NewError(i18n.Value("errors.handling.err_user_exists"), http.StatusBadRequest)
}
