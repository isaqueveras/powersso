package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPermission(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []string{})
}
