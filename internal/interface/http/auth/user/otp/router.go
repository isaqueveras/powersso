// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	gin "github.com/gin-gonic/gin"
	gopowersso "github.com/isaqueveras/go-powersso"
)

// Router is the router for the otp module.
func Router(r *gin.RouterGroup) {
	r.Use(gopowersso.SameUser())

	r.GET("qrcode", qrcode)
	r.POST("configure", configure)
	r.PUT("unconfigure", unconfigure)
}
