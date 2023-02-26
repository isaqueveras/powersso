// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import "github.com/gin-gonic/gin"

// Router is the router for the otp module.
func Router(r *gin.RouterGroup) {
	r.GET("qrcode", qrcode)
	r.POST("configure", configure)

	// TODO: Create route to unconfigure otp
	// TODO: Create route to check if OTP code is required when login
}
