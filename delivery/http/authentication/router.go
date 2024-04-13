// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package authentication

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/powersso/middleware"
)

// Router is the router for the auth module.
func Router(r *gin.RouterGroup) {
	r.POST("create_account", createAccount)
	r.POST("login", login)
	r.GET("login/steps", loginSteps)
	r.PUT("change_password", changePassword)
}

// RouterAuthorization is the router for the auth module.
func RouterAuthorization(r *gin.RouterGroup) {
	r.DELETE("logout", logout)

	user := r.Group("user/:user_id")
	user.PUT("disable", disable)

	otp := user.Group("otp")
	otp.Use(middleware.Yourself())

	otp.GET("qrcode", qrcode)
	otp.POST("configure", configure)
	otp.PUT("unconfigure", unconfigure)
}
