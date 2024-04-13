// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package authentication

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/powersso/application/authentication"
	domain "github.com/isaqueveras/powersso/domain/authentication"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

const sucessUserID = "9ec1b2a7-665c-47a7-b180-54f11f8a6122"

func TestHandlerAuth(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	router *gin.Engine

	suite.Suite
}

func (a *testSuite) SetupSuite() {
	var handleUserLog = func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Set("UID", sucessUserID)
			ctx.Set("SESSION", jwt.MapClaims{
				"SessionID": "",
				"UserID":    sucessUserID,
				"UserLevel": string(domain.AdminLevel),
				"FirstName": "Janekin",
			})
		}
	}

	a.router = gin.New()
	a.router.Use(middleware.RequestIdentifier(), handleUserLog())
	Router(a.router.Group("v1/auth"))
	RouterAuthorization(a.router.Group("v1/auth"))
	RouterAuthorization(a.router.Group("v1/auth/user/:user_id/otp"))
}

func (a *testSuite) TestShouldCreateUser() {
	monkey.Patch(authentication.CreateAccount, func(_ context.Context, _ *domain.CreateAccount) (*string, error) {
		return utils.Pointer(""), nil
	})
	defer monkey.Unpatch(authentication.CreateAccount)

	data, err := json.Marshal(map[string]interface{}{
		"first_name": "any_first_name",
		"last_name":  "any_last_name",
		"email":      "any@email.com",
		"password":   "any_password",
	})
	a.Assert().Nil(err, oops.Err(err))

	req := httptest.NewRequest(http.MethodPost, "/v1/auth/create_account", bytes.NewBuffer(data))
	w := httptest.NewRecorder()

	a.router.ServeHTTP(w, req)
	a.Assert().Equal(http.StatusCreated, w.Code)
}

func (t *testSuite) TestShouldGetUrlQrCode() {
	t.Run("Success", func() {
		monkey.Patch(authentication.GetQRCode2FA, func(_ context.Context, _ *uuid.UUID) (*string, error) {
			return nil, nil
		})
		defer monkey.Unpatch(authentication.GetQRCode2FA)

		req := httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+sucessUserID+"/otp/qrcode", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusOK, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(authentication.GetQRCode2FA, func(_ context.Context, _ *uuid.UUID) (*string, error) {
			return nil, nil
		})
		defer monkey.Unpatch(authentication.GetQRCode2FA)

		req := httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+uuid.New().String()+"/otp/qrcode", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}

func (t *testSuite) TestShouldConfigure() {
	t.Run("Success", func() {
		monkey.Patch(authentication.Configure2FA, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(authentication.Configure2FA)

		req := httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+sucessUserID+"/otp/configure", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusCreated, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(authentication.Configure2FA, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(authentication.Configure2FA)

		req := httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+uuid.New().String()+"/otp/configure", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}

func (t *testSuite) TestShouldUnconfigure() {
	t.Run("Success", func() {
		monkey.Patch(authentication.Unconfigure2FA, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(authentication.Unconfigure2FA)

		var (
			req = httptest.NewRequest(http.MethodPut, "/v1/auth/user/"+sucessUserID+"/otp/unconfigure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusCreated, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(authentication.Unconfigure2FA, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(authentication.Unconfigure2FA)

		var (
			req = httptest.NewRequest(http.MethodPut, "/v1/auth/user/"+uuid.New().String()+"/otp/unconfigure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}
