package box

import "github.com/gin-gonic/gin"

// Router ...
func Router(r *gin.RouterGroup) {
	r.GET("/my", getMyBoxes)
}
