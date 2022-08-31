// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tokens

import (
	"github.com/golang-jwt/jwt/v4"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/domain/user"
	"github.com/isaqueveras/power-sso/pkg/security"
)

// NewUserAuthToken generates and returns a new user authentication token.
func NewUserAuthToken(cfg *config.Config, user *user.User, sessionID *string) (string, error) {
	return security.NewToken(jwt.MapClaims{
		"session_id": sessionID,
		"user_id":    user.ID,
		"user_level": user.UserType,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	}, (cfg.UserAuthToken.SecretKey), cfg.UserAuthToken.Duration)
}
