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
)

func TestHandlerOTPInterface(t *testing.T) {
	suite.Run(t, new(otpHandlerSuite))
}

type otpHandlerSuite struct {
	router *gin.Engine

	suite.Suite
}

func (o *otpHandlerSuite) SetupSuite() {
	config.LoadConfig("../../../../../../")

	o.router = gin.New()
	Router(o.router.Group("v1/auth/user/:user_uuid/otp"))
}
func (o *otpHandlerSuite) TestShouldGetUrlQrCode() {
	monkey.Patch(otp.GetQRCode, func(_ context.Context, _ *uuid.UUID) (*otp.QRCodeResponse, error) {
		return &otp.QRCodeResponse{}, nil
	})
	defer monkey.Unpatch(otp.GetQRCode)

	var (
		req = httptest.NewRequest(http.MethodGet, "/v1/auth/user/"+uuid.New().String()+"/otp/qrcode", nil)
		w   = httptest.NewRecorder()
	)

	o.router.ServeHTTP(w, req)
	o.Assert().Equal(http.StatusOK, w.Code)
}
