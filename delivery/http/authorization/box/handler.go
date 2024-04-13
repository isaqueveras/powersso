// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package box

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/powersso/application/authorization/box"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

func getMyBoxes(ctx *gin.Context) {
	uid := middleware.GetSession(ctx).UserID

	boxes, err := box.GetMyBox(ctx, utils.Pointer(uid))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, boxes)
}
