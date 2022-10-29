// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"github.com/gin-gonic/gin"
	gopowersso "github.com/isaqueveras/go-powersso"
)

// RouterAuthorization is the router for the project module.
func RouterAuthorization(r *gin.RouterGroup) {
	r.POST("create", gopowersso.OnlyAdmin(), create)
}
