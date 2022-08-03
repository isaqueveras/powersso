// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"errors"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// GinZap add a custom middleware from zap
func GinZap(logger *zap.Logger, cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		c.Next()
		latency := time.Since(t1)

		fields := []zap.Field{
			zap.String("version", Version),
			zap.Time("date", time.Now()),
			zap.Int("status_code", c.Writer.Status()),
			zap.String("client_ip", c.ClientIP()),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Any("query", c.Copy().Request.URL.Query()),
			zap.Float64("latency", float64(latency)/float64(time.Millisecond)),
			zap.String("client_user_agent", c.Request.UserAgent()),
			zap.String("handler", strings.Join(strings.Split(strings.Replace(c.HandlerName(), cfg.Server.PermissionBase, "", 1), "/")[1:], "/")),
		}

		if requestId, ok := c.Get("RID"); ok {
			fields = append(fields, zap.Any("RID", requestId))
		}

		isError := false
		if errValue, set := c.Keys["error"]; set {
			var (
				err = errValue.(error)
				e   *oops.Error
			)

			if errors.As(err, &e) {
				fields = append(fields, []zap.Field{
					zap.Int("error_code", e.Code),
					zap.String("error", e.Error()),
					zap.String("cause", e.Err.Error()),
					zap.Strings("trace", e.Trace),
				}...)
				isError = true
			}
		}

		if isError {
			logger.Error("request handling failed", fields...)
		} else {
			logger.Info("requisition handled", fields...)
		}
	}
}

// RecoveryWithZap recovery middleware implementation with zap
func RecoveryWithZap(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return ginzap.RecoveryWithZap(logger, stack)
}
