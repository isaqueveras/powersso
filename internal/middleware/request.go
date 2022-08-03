// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestIdentifier inject a UUIDv4 in all contexts for better tracking
func RequestIdentifier() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("RID", uuid.New().String())
	}
}
