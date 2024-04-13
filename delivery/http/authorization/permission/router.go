// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package permission

import "github.com/gin-gonic/gin"

// Router is the router for the permission module.
func Router(r *gin.RouterGroup) {
	r.GET(":organization_id", get)
	r.POST("", create)
}
