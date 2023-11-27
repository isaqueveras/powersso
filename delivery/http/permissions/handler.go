package permissions

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/application/permission"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// @Router /v1/permission [get]
func getPermission(ctx *gin.Context) {
	pid, err := uuid.Parse(ctx.Param("pid"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	permissions, err := permission.GetPermissions(ctx, utils.Pointer(pid))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, permissions)
}
