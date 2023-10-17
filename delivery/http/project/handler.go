// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/powersso/application/project"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// @Router /v1/project/create [post]
func create(ctx *gin.Context) {
	var (
		input   = new(project.CreateProjectReq)
		session = middleware.GetSession(ctx)
		err     error
	)

	if err = ctx.ShouldBindJSON(&input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err = input.Validate(); err != nil {
		oops.Handling(ctx, err)
		return
	}

	input.CreatedByID = &session.UserID
	if err = project.Create(ctx, input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}
