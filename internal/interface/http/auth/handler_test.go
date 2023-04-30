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
	"github.com/stretchr/testify/suite"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/application/auth"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

func TestHandlerAuth(t *testing.T) {
	suite.Run(t, new(testSuite))
}

type testSuite struct {
	router *gin.Engine

	suite.Suite
}

func (a *testSuite) SetupSuite() {
	config.LoadConfig()

	a.router = gin.New()
	Router(a.router.Group("v1/auth"))
	RouterAuthorization(a.router.Group("v1/auth"))
}

func (a *testSuite) TestShouldCreateUser() {
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

func (t *testSuite) TestLoginSteps() {
	t.Run("UserFound", func() {
		monkey.Patch(auth.LoginSteps, func(ctx context.Context, email *string) (res *auth.StepsResponse, err error) {
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
		monkey.Patch(auth.LoginSteps, func(ctx context.Context, email *string) (res *auth.StepsResponse, err error) {
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
