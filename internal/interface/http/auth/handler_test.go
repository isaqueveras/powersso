// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bou.ke/monkey"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/application/auth"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

func TestHandlerAuthInterface(t *testing.T) {
	suite.Run(t, new(authHandlerSuite))
}

type authHandlerSuite struct {
	router *gin.Engine

	suite.Suite
}

func (a *authHandlerSuite) SetupSuite() {
	config.LoadConfig("../../../../")

	a.router = gin.New()
	Router(a.router.Group("v1/auth"))
	RouterAuthorization(a.router.Group("v1/auth"))
}

func (a *authHandlerSuite) TestShouldCreateUser() {
	monkey.Patch(auth.Register, func(_ context.Context, _ *auth.RegisterRequest) error {
		return nil
	})
	defer monkey.Unpatch(auth.Register)

	data, err := json.Marshal(map[string]interface{}{
		"first_name":   "any_first_name",
		"last_name":    "any_last_name",
		"email":        "any@email.com",
		"password":     "any_password",
		"phone_number": "any_phone_number",
		"address":      "any_address",
		"city":         "any_city",
		"country":      "any_country",
		"postcode":     55,
	})
	a.Assert().Nil(err, oops.Err(err))

	var (
		req = httptest.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewBuffer(data))
		w   = httptest.NewRecorder()
	)

	a.router.ServeHTTP(w, req)
	a.Assert().Equal(http.StatusCreated, w.Code)
}
