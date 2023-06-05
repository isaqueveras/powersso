// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/internal/utils"
	"github.com/isaqueveras/powersso/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

// Flag set the data type to flag the user
type Flag int64

const (
	// FlagOTPEnable defines that the user has OTP enabled
	FlagOTPEnable Flag = iota + 1
	// FlagOTPSetup defines that the user has OTP configured
	FlagOTPSetup
	// FlagEnabledAccount defines that the user has already activated his account
	FlagEnabledAccount
)

// Level set data type to user level
type Level string

const (
	// UserLevel is the user role
	UserLevel Level = "user"
	// AdminLevel is the admin role
	AdminLevel Level = "admin"
	// IntegrationLevel is the integration role
	IntegrationLevel Level = "integration"
)

const (
	// CostHashPasswordProduction is the cost of hashing password in production
	CostHashPasswordProduction int = 14
	// CostHashPasswordDevelopment is the cost of hashing the password in development mode
	CostHashPasswordDevelopment int = 1
)

// Register model the data to register user in the database
type Register struct {
	FirstName *string `sql:"first_name" json:"first_name"`
	LastName  *string `sql:"last_name" json:"last_name"`
	Email     *string `sql:"email" json:"email"`
	Password  *string `sql:"password" json:"password"`
	Role      *Flag   `sql:"roles" json:"-"`
	About     *string `sql:"about" json:"about"`
	Avatar    *string `sql:"avatar" json:"avatar"`
	TokenKey  *string `sql:"token_key" json:"token_key"`
}

// Prepare prepare data for registration
func (rr *Register) Prepare() (err error) {
	rr.Email = utils.GetStringPointer(strings.ToLower(strings.TrimSpace(*rr.Email)))
	rr.Password = utils.GetStringPointer(strings.TrimSpace(*rr.Password))

	if err = rr.GeneratePassword(); err != nil {
		return err
	}

	return
}

// RefreshTokenKey generates and sets new random token key.
// >> invalidate previously issued tokens
func (rr *Register) RefreshTokenKey() {
	rr.TokenKey = new(string)
	rr.TokenKey = utils.GetStringPointer(security.RandomString(50))
}

// GeneratePassword hash user password with bcrypt
func (rr *Register) GeneratePassword() error {
	rr.RefreshTokenKey()

	cost := CostHashPasswordDevelopment
	if config.Get().Server.IsModeProduction() {
		cost = CostHashPasswordProduction
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*rr.TokenKey+*rr.Password), cost)
	if err != nil {
		return err
	}

	rr.Password = utils.GetStringPointer(string(hashedPassword))
	return nil
}

// SanitizePassword sanitize user password
func (rr *Register) SanitizePassword() {
	rr.Password = nil
}

// ActivateAccountToken model the data to activate user account
type ActivateAccountToken struct {
	ID        *uuid.UUID
	UserID    *uuid.UUID
	Used      *bool
	Valid     *bool
	ExpiresAt *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// IsValid check if the token is valid
func (a *ActivateAccountToken) IsValid() bool {
	return (a.Used != nil && *a.Used) || (a.Valid != nil && !*a.Valid)
}

// Steps contains login steps
type Steps struct {
	Name *string
	OTP  *bool
}

type User struct {
	ID        *uuid.UUID `sql:"id"`
	Email     *string    `sql:"email"`
	Password  *string    `sql:"password" json:"-"`
	FirstName *string    `sql:"first_name"`
	LastName  *string    `sql:"last_name"`
	About     *string    `sql:"about"`
	Flag      *Flag      `sql:"flag"`
	Level     *Level     `sql:"level"`
	Avatar    *string    `sql:"avatar"`
	Blocked   *bool      `sql:"attempts >= 3 AND (U.last_failure + '30 minutes') >= NOW()"`
	Key       *string    `sql:"key"`
	Active    *bool      `sql:"active"`
	OTPToken  *string    `sql:"otp_token"`
	CreatedBy *uuid.UUID `sql:"created_by"`
	CreatedAt *time.Time `sql:"created_at"`
	LastLogin *time.Time `sql:"last_login"`
}

// HasFlag return 'true' if has flag
func (u *User) HasFlag(flag Flag) bool {
	return u.Flag != nil && *u.Flag&flag != 0
}

// IsActive check if the user has their account activated
func (u *User) IsActive() bool {
	return u.Active != nil && *u.Active
}

// IsBlocked check if the user has the account temporarily blocked
func (u *User) IsBlocked() bool {
	return u.Blocked != nil && *u.Blocked
}

func (u *User) OTPConfigured() bool {
	enabled := u.Flag != nil && *u.Flag&FlagOTPEnable != 0
	setup := u.Flag != nil && *u.Flag&FlagOTPSetup != 0
	return enabled && setup
}

// QRCode wraps the data to return the qr code url
type QRCode struct {
	Url *string `json:"url,omitempty"`
}

// Login models the data for the user to log in with their account
type Login struct {
	Email     *string `json:"email" binding:"required,lte=60,email"`
	Password  *string `json:"password" binding:"required,gte=6"`
	OTP       *string `json:"otp,omitempty"`
	ClientIP  *string `json:"-"`
	UserAgent *string `json:"-"`
}

// ComparePasswords compare user password and payload
func (l *Login) ComparePasswords(passw, tokenKey *string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(*passw), []byte(*tokenKey+*l.Password)); err != nil {
		return ErrEmailOrPasswordIsNotValid()
	}
	l.SanitizePassword()
	return nil
}

// SanitizePassword sanitize user password
func (l *Login) SanitizePassword() {
	l.Password = nil
}

// Validate prepare data for login
func (l *Login) Validate() {
	if l.ClientIP != nil && *l.ClientIP == "" {
		l.ClientIP = utils.GetStringPointer("0.0.0.0")
	}

	if l.UserAgent != nil && *l.UserAgent == "" {
		l.UserAgent = utils.GetStringPointer("Unknown")
	}
}

// Session models the data of a user session
type Session struct {
	SessionID *uuid.UUID     `json:"session_id,omitempty"`
	UserID    *uuid.UUID     `json:"user_id,omitempty"`
	Email     *string        `json:"email,omitempty"`
	FirstName *string        `json:"first_name,omitempty"`
	LastName  *string        `json:"last_name,omitempty"`
	About     *string        `json:"about,omitempty"`
	Level     *Level         `json:"level,omitempty"`
	Token     *string        `json:"token,omitempty"`
	CreatedAt *time.Time     `json:"created_at,omitempty"`
	ExpiresAt *time.Time     `json:"expires_at,omitempty"`
	RawData   map[string]any `json:"data,omitempty"`
}
