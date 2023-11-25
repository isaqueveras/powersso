// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/isaqueveras/powersso/application/project"
	"github.com/isaqueveras/powersso/middleware"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// @Router /v1/project/create [post]
func newProject(ctx *gin.Context) {
	var (
		input   = new(project.NewProject)
		session = middleware.GetSession(ctx)
	)

	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	input.Slug = utils.Pointer(slug.Make(*input.Name))
	input.CreatedByID = &session.UserID

	if err := project.CreateNewProject(ctx, input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}
