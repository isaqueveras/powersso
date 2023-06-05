// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/powersso/internal/application/auth"
	domain "github.com/isaqueveras/powersso/internal/domain/auth"
	"github.com/isaqueveras/powersso/internal/middleware"
	"github.com/isaqueveras/powersso/pkg/oops"
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
		return func(ctx *gin.Context) { ctx.Set("UID", sucessUserID) }
	}

	a.router = gin.New()
	a.router.Use(middleware.RequestIdentifier(), handleUserLog())
	Router(a.router.Group("v1/auth"))
	RouterAuthorization(a.router.Group("v1/auth"))
	RouterAuthorization(a.router.Group("v1/auth/user/:user_uuid/otp"))
}

func (a *testSuite) TestShouldCreateUser() {
	monkey.Patch(auth.Register, func(_ context.Context, _ *domain.Register) error {
		return nil
	})
	defer monkey.Unpatch(auth.Register)

	data, err := json.Marshal(map[string]interface{}{
		"first_name": "any_first_name",
		"last_name":  "any_last_name",
		"email":      "any@email.com",
		"password":   "any_password",
	})
	a.Assert().Nil(err, oops.Err(err))

	var (
		req = httptest.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewBuffer(data))
		w   = httptest.NewRecorder()
	)

	a.router.ServeHTTP(w, req)
	a.Assert().Equal(http.StatusCreated, w.Code)
}

func (t *testSuite) TestLoginSteps() {
	t.Run("UserFound", func() {
		monkey.Patch(auth.LoginSteps, func(ctx context.Context, email *string) (res *domain.Steps, err error) {
			return nil, nil
		})
		defer monkey.Unpatch(auth.LoginSteps)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/login/steps?email=luiz@bonfa.com", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusOK, w.Code)
	})

	t.Run("UserNotFound", func() {
		monkey.Patch(auth.LoginSteps, func(ctx context.Context, email *string) (*domain.Steps, error) {
			return nil, sql.ErrNoRows
		})
		defer monkey.Unpatch(auth.LoginSteps)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/login/steps", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusNotFound, w.Code)
	})
}

func (t *testSuite) TestShouldGetUrlQrCode() {
	t.Run("Success", func() {
		monkey.Patch(auth.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*domain.QRCode, error) {
			return &domain.QRCode{}, nil
		})
		defer monkey.Unpatch(auth.GetQrCode)

		req := httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+sucessUserID+"/otp/qrcode", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusOK, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(auth.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*domain.QRCode, error) {
			return &domain.QRCode{}, nil
		})
		defer monkey.Unpatch(auth.GetQrCode)

		req := httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+uuid.New().String()+"/otp/qrcode", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}

func (t *testSuite) TestShouldConfigure() {
	t.Run("Success", func() {
		monkey.Patch(auth.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(auth.Configure)

		req := httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+sucessUserID+"/otp/configure", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusCreated, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(auth.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(auth.Configure)

		req := httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+uuid.New().String()+"/otp/configure", nil)
		w := httptest.NewRecorder()

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}

func (t *testSuite) TestShouldUnconfigure() {
	t.Run("Success", func() {
		monkey.Patch(auth.Unconfigure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(auth.Unconfigure)

		var (
			req = httptest.NewRequest(http.MethodPut, "/v1/auth/user/"+sucessUserID+"/otp/unconfigure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusCreated, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(auth.Unconfigure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(auth.Unconfigure)

		var (
			req = httptest.NewRequest(http.MethodPut, "/v1/auth/user/"+uuid.New().String()+"/otp/unconfigure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}
