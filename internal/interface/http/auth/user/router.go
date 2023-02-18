// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/power-sso/internal/interface/http/auth/user/otp"
)

// RouterWithUUID is the router for the user module.
func RouterWithUUID(r *gin.RouterGroup) {
	r.PUT("disable", disable)

	otp.Router(r.Group("otp"))
}
