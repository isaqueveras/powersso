// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/power-sso/internal/interface/http/auth/user"
)

// Router is the router for the auth module.
func Router(r *gin.RouterGroup) {
	r.POST("activation/:token", activation)
	r.POST("register", register)
	r.POST("login", login)
	r.GET("login/steps", loginSteps)
}

// RouterAuthorization is the router for the auth module.
func RouterAuthorization(r *gin.RouterGroup) {
	r.DELETE("logout", logout)

	user.RouterWithUUID(r.Group("user/:user_uuid"))
}
