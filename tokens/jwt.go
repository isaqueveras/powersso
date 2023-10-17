// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// NewToken generates and returns new HS256 signed JWT token.
func NewToken(payload jwt.MapClaims, key string, duration int64) (string, error) {
	var (
		seconds = time.Duration(duration) * time.Second
		claims  = jwt.MapClaims{"exp": time.Now().Add(seconds).Unix()}
	)

	for key, value := range payload {
		claims[key] = value
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))
}

// ParseJWT verifies and parses JWT token and returns its claims.
func ParseJWT(token string, keys []string) jwt.MapClaims {
	parser := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"}))
	for _, key := range keys {
		parsed, _ := parser.Parse(token, func(t *jwt.Token) (interface{}, error) { return []byte(key), nil })
		if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
			return claims
		}
	}
	return nil
}
