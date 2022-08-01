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
