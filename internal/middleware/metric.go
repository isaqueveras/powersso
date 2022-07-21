package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/power-sso/config"
	"go.uber.org/zap"
)

// GinZap adiciona um middleware customizado do zap
func GinZap(logger *zap.Logger, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		c.Next()
		latency := time.Since(t1)

		fields := []zap.Field{
			zap.Time("date", time.Now()),
			zap.Int("status_code", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Float64("latency", float64(latency)/float64(time.Millisecond)),
			zap.String("client_user_agent", c.Request.UserAgent()),
			zap.String("handler", strings.Join(strings.Split(strings.Replace(c.HandlerName(), cfg.Server.PermissionBase, "", 1), "/")[1:], "/")),
		}

		logger.Info("requisition handled", fields...)
	}
}
