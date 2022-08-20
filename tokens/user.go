// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tokens

import (
	"github.com/golang-jwt/jwt/v4"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/pkg/security"
)

// NewUserAuthToken generates and returns a new user authentication token.
func NewUserAuthToken(cfg *config.Config, userID, tokenKey *string) (string, error) {
	return security.NewToken(jwt.MapClaims{"user_id": userID, "type": "user"}, (*tokenKey + cfg.UserAuthToken.SecretKey), cfg.UserAuthToken.Duration)
}
