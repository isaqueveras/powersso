// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tokens

import (
	"log"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/utils"
)

// NewAuthToken generates and returns a new authentication token
func NewAuthToken(user *auth.User, sessionID *uuid.UUID) (*string, error) {
	claims := jwt.MapClaims{
		"session_id": sessionID,
		"user_id":    user.ID,
		"user_level": user.Level,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	}

	log.Println(user.GetUserLevel(&config.Get().SecretsTokens), config.Get().SecretsDuration)

	token, err := utils.NewToken(claims, user.GetUserLevel(&config.Get().SecretsTokens), config.Get().SecretsDuration)
	return utils.Pointer(token), err
}
