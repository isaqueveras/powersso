// Copyright (c) 2024 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/powersso/application/authorization/organization"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// @Router /v1/room/create [post]
func create(ctx *gin.Context) {
	var (
		in      = new(organization.Organization)
		session = middleware.GetSession(ctx)
	)

	if err := ctx.ShouldBindJSON(&in); err != nil {
		oops.Handling(ctx, err)
		return
	}

	in.CreatedByID = utils.Pointer(session.UserID)
	if err := organization.Create(ctx, in); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}
