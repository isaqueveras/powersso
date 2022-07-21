package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Version contains version information
var Version string = ""

// VersionInfo add a version header to the request
func VersionInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Version == "" {
			Version = strconv.FormatInt(time.Now().Unix(), 10)
		}
		c.Writer.Header().Set("Application-Version", Version)
	}
}
