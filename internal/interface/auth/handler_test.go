// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/interface/auth"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
)

// TestIntegrationAuth is a test for the auth package.
func TestIntegrationAuth(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config.LoadConfig("../../../")

	var (
		mock sqlmock.Sqlmock
		err  error
	)

	if mock, err = postgres.OpenConnectionForTesting(); err != nil {
		t.Fatal(err)
	}
	defer postgres.CloseConnections()

	router := gin.Default()
	auth.Router(router.Group("v1/auth"))
	auth.RouterAuthorization(router.Group("v1/auth"))

	t.Run("Register", func(t *testing.T) {
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
		assert.Equal(t, err, nil)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT COUNT(id) > 0 FROM users WHERE email = $1`)).
			WithArgs("any@email.com").
			WillReturnRows(sqlmock.NewRows([]string{"exist"}).AddRow(false))

		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO users (first_name,last_name,email,password,roles,phone_number,address,city,country,postcode,token_key) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING "id"`)).
			WithArgs("any_first_name", "any_last_name", "any@email.com", "any_password", "{read:activation_token}", "any_phone_number", "any_address", "any_city", "any_country", 55, "token_key_testing_any_password").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

		mock.ExpectCommit()

		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(data))
		assert.Equal(t, err, nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, 201, w.Code)
	})
}
