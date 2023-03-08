// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/application/auth/user/otp"
	"github.com/isaqueveras/power-sso/internal/middleware"
)

const sucessUserID = "9ec1b2a7-665c-47a7-b180-54f11f8a6122"

func TestHandlerOTPInterface(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	router *gin.Engine

	suite.Suite
}

func (o *testSuite) SetupSuite() {
	config.LoadConfig("../../../../../../")
	var handleUserLog = func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Set("UID", sucessUserID)
		}
	}

	o.router = gin.New()
	o.router.Use(middleware.RequestIdentifier(), handleUserLog())
	Router(o.router.Group("v1/auth/user/:user_uuid/otp"))
}

func (t *testSuite) TestShouldGetUrlQrCode() {
	t.Run("Success", func() {
		monkey.Patch(otp.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*otp.QRCodeResponse, error) {
			return &otp.QRCodeResponse{}, nil
		})
		defer monkey.Unpatch(otp.GetQrCode)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+sucessUserID+"/otp/qrcode", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusOK, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(otp.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*otp.QRCodeResponse, error) {
			return &otp.QRCodeResponse{}, nil
		})
		defer monkey.Unpatch(otp.GetQrCode)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+uuid.New().String()+"/otp/qrcode", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}

func (t *testSuite) TestShouldConfigure() {
	t.Run("Success", func() {
		monkey.Patch(otp.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Configure)

		var (
			req = httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+sucessUserID+"/otp/configure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusCreated, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(otp.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Configure)

		var (
			req = httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+uuid.New().String()+"/otp/configure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}

func (t *testSuite) TestShouldUnconfigure() {
	t.Run("Success", func() {
		monkey.Patch(otp.Unconfigure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Unconfigure)

		var (
			req = httptest.NewRequest(http.MethodPut, "/v1/auth/user/"+sucessUserID+"/otp/unconfigure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusCreated, w.Code)
	})

	t.Run("Error::FetchAnotherUserURL", func() {
		monkey.Patch(otp.Unconfigure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Unconfigure)

		var (
			req = httptest.NewRequest(http.MethodPut, "/v1/auth/user/"+uuid.New().String()+"/otp/unconfigure", nil)
			w   = httptest.NewRecorder()
		)

		t.router.ServeHTTP(w, req)
		t.Assert().Equal(http.StatusForbidden, w.Code)
	})
}
