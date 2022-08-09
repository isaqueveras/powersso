// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package tokens

import (
	"github.com/golang-jwt/jwt/v4"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/pkg/security"
)

// NewUserVerifyToken generates and returns a new user verification token.
func NewUserVerifyToken(cfg *config.Config, email, userTokenKey *string) (string, error) {
	payload := jwt.MapClaims{
		"type":  "user",
		"email": email,
	}

	return security.NewToken(payload, (*userTokenKey + cfg.UserVerificationToken.Secret), cfg.UserVerificationToken.Duration)
}
