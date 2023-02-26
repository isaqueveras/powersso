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
	gopowersso "github.com/isaqueveras/go-powersso"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/application/auth/user/otp"
	"github.com/isaqueveras/power-sso/internal/middleware"
	"github.com/isaqueveras/power-sso/pkg/logger"
)

func TestHandlerOTPInterface(t *testing.T) {
	suite.Run(t, new(otpHandlerSuite))
}

type otpHandlerSuite struct {
	router *gin.Engine
	cfg    *config.Config

	suite.Suite
}

func (o *otpHandlerSuite) SetupSuite() {
	config.LoadConfig("../../../../../../")
	o.cfg = config.Get()

	logg := logger.NewLogger(o.cfg)
	logg.InitLogger()

	o.router = gin.New()
	o.router.Use(middleware.RequestIdentifier())
	Router(o.router.Group("v1/auth/user/:user_uuid/otp"))
}

func (o *otpHandlerSuite) TestShouldGetUrlQrCode() {
	o.Run("Success", func() {
		monkey.Patch(gopowersso.SameUserRequest, func(_ *gin.Context, _ string) bool {
			return true
		})
		defer monkey.Unpatch(gopowersso.SameUserRequest)

		monkey.Patch(otp.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*otp.QRCodeResponse, error) {
			return &otp.QRCodeResponse{}, nil
		})
		defer monkey.Unpatch(otp.GetQrCode)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+uuid.New().String()+"/otp/qrcode", nil)
			w   = httptest.NewRecorder()
		)

		o.router.ServeHTTP(w, req)
		o.Assert().Equal(http.StatusOK, w.Code)
	})

	o.Run("Error > Fetch another user's URL", func() {
		monkey.Patch(otp.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*otp.QRCodeResponse, error) {
			return &otp.QRCodeResponse{}, nil
		})
		defer monkey.Unpatch(otp.GetQrCode)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+uuid.New().String()+"/otp/qrcode", nil)
			w   = httptest.NewRecorder()
		)

		o.router.ServeHTTP(w, req)
		o.Assert().Equal(http.StatusBadRequest, w.Code)
	})

	o.Run("Error > Invalid UUID format", func() {
		monkey.Patch(otp.GetQrCode, func(_ context.Context, _ *uuid.UUID) (*otp.QRCodeResponse, error) {
			return &otp.QRCodeResponse{}, nil
		})
		defer monkey.Unpatch(otp.GetQrCode)

		var (
			req = httptest.NewRequest(http.MethodGet, "/v1/auth/user/213213-23234-2341sdf-234d/otp/qrcode", nil)
			w   = httptest.NewRecorder()
		)

		o.router.ServeHTTP(w, req)
		o.Assert().Equal(http.StatusBadRequest, w.Code)
	})
}

func (o *otpHandlerSuite) TestShouldConfigure() {
	o.Run("Success", func() {
		monkey.Patch(gopowersso.SameUserRequest, func(_ *gin.Context, _ string) bool {
			return true
		})
		defer monkey.Unpatch(gopowersso.SameUserRequest)

		monkey.Patch(otp.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Configure)

		var (
			req = httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+uuid.New().String()+"/otp/configure", nil)
			w   = httptest.NewRecorder()
		)

		o.router.ServeHTTP(w, req)
		o.Assert().Equal(http.StatusCreated, w.Code)
	})

	o.Run("Error > Fetch another user's URL", func() {
		monkey.Patch(otp.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Configure)

		var (
			req = httptest.NewRequest(http.MethodPost, "/v1/auth/user/"+uuid.New().String()+"/otp/configure", nil)
			w   = httptest.NewRecorder()
		)

		o.router.ServeHTTP(w, req)
		o.Assert().Equal(http.StatusBadRequest, w.Code)
	})

	o.Run("Error > Invalid UUID format", func() {
		monkey.Patch(otp.Configure, func(_ context.Context, _ *uuid.UUID) error {
			return nil
		})
		defer monkey.Unpatch(otp.Configure)

		var (
			req = httptest.NewRequest(http.MethodPost, "/v1/auth/user/213213-23234-2341sdf-234d/otp/configure", nil)
			w   = httptest.NewRecorder()
		)

		o.router.ServeHTTP(w, req)
		o.Assert().Equal(http.StatusBadRequest, w.Code)
	})
}
