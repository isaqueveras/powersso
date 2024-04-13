// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package box

import "github.com/gin-gonic/gin"

// Router ...
func Router(r *gin.RouterGroup) {
	r.GET("/my", getMyBoxes)
}
