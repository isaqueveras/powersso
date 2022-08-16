// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import "github.com/gin-gonic/gin"

func Router(r *gin.RouterGroup) {
	r.POST("/register", register)
	r.POST("/activation/:token", activation)
	r.POST("/login", login)
}
