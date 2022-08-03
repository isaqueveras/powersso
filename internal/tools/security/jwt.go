// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// NewToken generates and returns new HS256 signed JWT token.
func NewToken(payload jwt.MapClaims, signingKey string, secondsDuration int64) (string, error) {
	var (
		seconds = time.Duration(secondsDuration) * time.Second
		claims  = jwt.MapClaims{
			"exp": time.Now().Add(seconds).Unix(),
		}
	)

	for key, value := range payload {
		claims[key] = value
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(signingKey))
}

// ParseJWT verifies and parses JWT token and returns its claims.
func ParseJWT(token, verificationKey string) (jwt.MapClaims, error) {
	var (
		parser      = jwt.NewParser(jwt.WithValidMethods([]string{"HS256"}))
		parsedToken *jwt.Token
		err         error
	)

	if parsedToken, err = parser.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(verificationKey), nil
	}); err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("unable to parse token")
}

// ParseUnverifiedJWT parses JWT token and returns its claims
// but does not verify the signature.
func ParseUnverifiedJWT(token string) (jwt.MapClaims, error) {
	var parser *jwt.Parser = &jwt.Parser{}
	var claims = jwt.MapClaims{}
	var err error

	if _, _, err = parser.ParseUnverified(token, claims); err == nil {
		err = claims.Valid()
	}

	return claims, err
}
