// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package tokens_test

import (
	"testing"

	"github.com/golang-jwt/jwt/v4"
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/pkg/security"
	"github.com/isaqueveras/power-sso/tokens"
)

func TestNewUserVerifyToken(t *testing.T) {
	config.LoadConfig("../")

	var (
		email    = "test@email.com"
		tokenKey = security.RandomString(50)
		cfg      = config.Get()

		token string
		err   error
	)

	if token, err = tokens.NewUserVerifyToken(cfg, &email, &tokenKey); err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Fatal("token is nil")
	}

	var claims jwt.MapClaims
	if claims, err = security.ParseJWT(token, (tokenKey + cfg.UserVerificationToken.Secret)); err != nil {
		t.Fatal(err)
	}

	if claims["type"] != "user" {
		t.Fatal("type is not user")
	}

	if claims["email"] != email {
		t.Fatal("email is not equal")
	}
}
