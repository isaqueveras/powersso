// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package permission

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/isaqueveras/powersso/application/authorization/permission"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// @Router /v1/permission/{organization_id} [get]
func get(ctx *gin.Context) {
	pid, err := uuid.Parse(ctx.Param("organization_id"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	session := middleware.GetSession(ctx)

	var permissions []*string
	if permissions, err = permission.Get(ctx, utils.Pointer(session.UserID), utils.Pointer(pid)); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string][]*string{
		"permissions": permissions,
	})
}

// @Router /v1/permission [post]
func create(ctx *gin.Context) {
	input := &permission.Permission{}
	if err := ctx.ShouldBindJSON(input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	session := middleware.GetSession(ctx)
	input.CreatedByID = utils.Pointer(session.UserID)

	if err := permission.Create(ctx, input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
