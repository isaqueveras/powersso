package permissions

import "github.com/gin-gonic/gin"

// Router is the router for the permission module.
func Router(r *gin.RouterGroup) {
	r.GET("", getPermission)
}
