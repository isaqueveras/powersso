// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/domain/authentication"
	"github.com/isaqueveras/powersso/utils"
)

// NewAuthToken generates and returns a new authentication token
func NewAuthToken(user *authentication.User, sessionID *uuid.UUID) (*string, error) {
	claims := jwt.MapClaims{
		"SessionID": sessionID,
		"UserID":    user.ID,
		"UserLevel": user.Level,
		"FirstName": user.FirstName,
	}

	token, err := NewToken(claims, user.GetUserLevel(&config.Get().SecretsTokens), config.Get().SecretsDuration)
	return utils.Pointer(token), err
}
