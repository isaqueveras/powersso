// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/isaqueveras/power-sso/internal/domain/auth/roles"
	"github.com/isaqueveras/power-sso/pkg/security"
)

// RegisterRequest is the request payload for the register endpoint.
type RegisterRequest struct {
	FirstName   *string    `json:"first_name" binding:"required,lte=30"`
	LastName    *string    `json:"last_name" binding:"required,lte=30"`
	Email       *string    `json:"email,omitempty" binding:"omitempty,lte=60,email"`
	Password    *string    `json:"password,omitempty" binding:"omitempty,required,gte=6"`
	About       *string    `json:"about,omitempty" binding:"omitempty,lte=1024"`
	Avatar      *string    `json:"avatar,omitempty" binding:"omitempty,lte=512,url"`
	PhoneNumber *string    `json:"phone_number,omitempty" binding:"omitempty,lte=20"`
	Address     *string    `json:"address,omitempty" binding:"omitempty,lte=250"`
	City        *string    `json:"city,omitempty" binding:"omitempty,lte=24"`
	Country     *string    `json:"country,omitempty" binding:"omitempty,lte=24"`
	Gender      *string    `json:"gender,omitempty" binding:"omitempty,lte=10"`
	Postcode    *int       `json:"postcode,omitempty" binding:"omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty" binding:"omitempty,lte=10"`
	TokenKey    *string    `json:"token_key"`

	Roles *roles.Roles `json:"-"`
}

// Prepare prepare data for registration
func (rr *RegisterRequest) Prepare() (err error) {
	*rr.Email = strings.ToLower(strings.TrimSpace(*rr.Email))
	*rr.Password = strings.TrimSpace(*rr.Password)

	if err = rr.GeneratePassword(); err != nil {
		return err
	}

	if rr.PhoneNumber != nil {
		*rr.PhoneNumber = strings.TrimSpace(*rr.PhoneNumber)
	}

	return
}

// GeneratePassword hash user password with bcrypt
func (rr *RegisterRequest) GeneratePassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*rr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*rr.Password = string(hashedPassword)
	rr.RefreshTokenKey()

	return nil
}

// ComparePasswords compare user password and payload
func (rr *RegisterRequest) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(*rr.Password), []byte(password)); err != nil {
		return err
	}
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
	*rr.TokenKey = security.RandomString(50)
}

// LoginRequest is the request payload for the login endpoint.
type LoginRequest struct {
	Email    *string `json:"email" binding:"required,lte=60,email"`
	Password *string `json:"password" binding:"required,gte=6"`
}

// ComparePasswords compare user password and payload
func (lr *LoginRequest) ComparePasswords(passw *string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(*passw), []byte(*lr.Password)); err != nil {
		return ErrEmailOrPasswordIsNotValid()
	}
	return
}

// SanitizePassword sanitize user password
func (rr *LoginRequest) SanitizePassword() {
	rr.Password = nil
}
