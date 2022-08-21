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
	"testing"

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
	cfg := config.Get()

	if err := postgres.OpenConnections(cfg); err != nil {
		t.Fatal("Unable to open connections to database: ", err)
	}
	defer postgres.CloseConnections()

	router := gin.Default()
	auth.Router(router.Group("v1/auth"))

	t.Run("Register", func(t *testing.T) {
		data, err := json.Marshal(map[string]interface{}{
			"first_name":   "any_first_name",
			"last_name":    "any_last_name",
			"email":        "any@email.com",
			"password":     "any_password",
			"about":        "any_about",
			"phone_number": "any_phone_number",
			"address":      "any_address",
			"city":         "any_city",
			"country":      "any_country",
			"gender":       "any_gender",
			"postcode":     55,
			"birthday":     "2022-08-04T19:44:00-03:00",
		})
		assert.Equal(t, err, nil)
		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(data))
		assert.Equal(t, err, nil)

		router.ServeHTTP(w, req)
		assert.Equal(t, 201, w.Code)
	})
}
