// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tokens

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/utils"
)

// NewUserAuthToken generates and returns a new user authentication token.
func NewUserAuthToken(user *auth.User, sessionID *uuid.UUID) (*string, error) {
	claims := jwt.MapClaims{
		"session_id": sessionID,
		"user_id":    user.ID,
		"user_level": user.Level,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	}

	token, err := utils.NewToken(claims, (config.Get().UserAuthToken.SecretKey), config.Get().UserAuthToken.Duration)
	return utils.Pointer(token), err
}
