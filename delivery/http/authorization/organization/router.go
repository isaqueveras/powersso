// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package organization

import "github.com/gin-gonic/gin"

// Router is the router for the room module.
func Router(r *gin.RouterGroup) {
	r.POST("create", create)
}
