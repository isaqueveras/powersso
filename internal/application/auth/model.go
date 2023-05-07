// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/isaqueveras/power-sso/config"
	domain "github.com/isaqueveras/power-sso/internal/domain/auth"
	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	"github.com/isaqueveras/power-sso/internal/domain/auth/user"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/pkg/security"
)

type (
	// RegisterRequest is the request payload for the register endpoint.
	RegisterRequest struct {
		FirstName   *string      `json:"first_name" binding:"required,lte=30"`
		LastName    *string      `json:"last_name" binding:"required,lte=30"`
		Email       *string      `json:"email,omitempty" binding:"omitempty,required,lte=60,email"`
		Password    *string      `json:"password,omitempty" binding:"omitempty,required,gte=6"`
		About       *string      `json:"about,omitempty" binding:"omitempty,lte=1024"`
		Avatar      *string      `json:"avatar,omitempty" binding:"omitempty,lte=512,url"`
		PhoneNumber *string      `json:"phone_number,omitempty" binding:"omitempty,lte=20"`
		Address     *string      `json:"address,omitempty" binding:"omitempty,lte=250"`
		City        *string      `json:"city,omitempty" binding:"omitempty,lte=24"`
		Country     *string      `json:"country,omitempty" binding:"omitempty,lte=24"`
		Gender      *string      `json:"gender,omitempty" binding:"omitempty,lte=10"`
		Postcode    *int         `json:"postcode,omitempty" binding:"omitempty"`
		Birthday    *time.Time   `json:"birthday,omitempty" binding:"omitempty,lte=10"`
		TokenKey    *string      `json:"token_key"`
		Roles       *roles.Roles `json:"-"`
	}

	// SessionResponse define a session model output for presentation layer
	SessionResponse struct {
		SessionID   *string        `json:"session_id,omitempty"`
		Level       *user.Level    `json:"level,omitempty"`
		UserID      *string        `json:"user_id,omitempty"`
		Email       *string        `json:"email,omitempty"`
		FirstName   *string        `json:"first_name,omitempty"`
		LastName    *string        `json:"last_name,omitempty"`
		About       *string        `json:"about,omitempty"`
		AvatarURL   *string        `json:"avatar_url,omitempty"`
		PhoneNumber *string        `json:"phone_number,omitempty"`
		OTPEnabled  *bool          `json:"otp_enabled,omitempty"`
		OTPSetUp    *bool          `json:"otp_setup,omitempty"`
		Roles       []string       `json:"roles,omitempty"`
		Token       *string        `json:"token,omitempty"`
		RawData     map[string]any `json:"data,omitempty"`
		CreatedAt   *time.Time     `json:"created_at,omitempty"`
		ExpiresAt   *time.Time     `json:"expires_at,omitempty"`
	}

	// StepsResponse returns the data for login
	StepsResponse struct {
		Name *string `json:"name,omitempty"`
		OTP  *bool   `json:"otp,omitempty"`
	}
)

// Prepare prepare data for registration
func (rr *RegisterRequest) Prepare() (err error) {
	rr.Email = utils.GetStringPointer(strings.ToLower(strings.TrimSpace(*rr.Email)))
	rr.Password = utils.GetStringPointer(strings.TrimSpace(*rr.Password))

	if err = rr.GeneratePassword(); err != nil {
		return err
	}

	if rr.PhoneNumber != nil {
		rr.PhoneNumber = utils.GetStringPointer(strings.TrimSpace(*rr.PhoneNumber))
	}

	return
}

// GeneratePassword hash user password with bcrypt
func (rr *RegisterRequest) GeneratePassword() error {
	rr.RefreshTokenKey()

	cost := domain.CostHashPasswordDevelopment
	if config.Get().Server.IsModeProduction() {
		cost = domain.CostHashPasswordProduction
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*rr.TokenKey+*rr.Password), cost)
	if err != nil {
		return err
	}

	rr.Password = utils.GetStringPointer(string(hashedPassword))
	return nil
}

// SanitizePassword sanitize user password
func (rr *RegisterRequest) SanitizePassword() {
	rr.Password = nil
}

// RefreshTokenKey generates and sets new random token key.
// >> invalidate previously issued tokens
func (rr *RegisterRequest) RefreshTokenKey() {
	rr.TokenKey = new(string)
	rr.TokenKey = utils.GetStringPointer(security.RandomString(50))
}

// LoginRequest is the request payload for the login endpoint.
type LoginRequest struct {
	Email    *string `json:"email" binding:"required,lte=60,email"`
	Password *string `json:"password" binding:"required,gte=6"`
	OTP      *string `json:"otp,omitempty"`

	ClientIP  string `json:"-"`
	UserAgent string `json:"-"`
}

// Validate prepare data for login
func (lr *LoginRequest) Validate() {
	if lr.ClientIP == "" {
		lr.ClientIP = "0.0.0.0"
	}

	if lr.UserAgent == "" {
		lr.UserAgent = "Unknown"
	}
}

// ComparePasswords compare user password and payload
func (lr *LoginRequest) ComparePasswords(passw, tokenKey *string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(*passw), []byte(*tokenKey+*lr.Password)); err != nil {
		return domain.ErrEmailOrPasswordIsNotValid()
	}
	lr.SanitizePassword()
	return nil
}

// SanitizePassword sanitize user password
func (rr *LoginRequest) SanitizePassword() {
	rr.Password = nil
}
