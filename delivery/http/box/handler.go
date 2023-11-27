package box

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isaqueveras/powersso/application/box"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

func getMyBoxes(ctx *gin.Context) {
	uid, err := uuid.Parse(middleware.GetSession(ctx).UserID)
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	boxes, err := box.GetMyBox(ctx, utils.Pointer(uid))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, boxes)
}
