// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gopowersso "github.com/isaqueveras/go-powersso"

	"github.com/isaqueveras/powersso/internal/application/project"
	"github.com/isaqueveras/powersso/internal/utils"
	"github.com/isaqueveras/powersso/pkg/oops"
)

// create godoc
// @Summary Register a projet
// @Description Register a project including several users to the project
// @Tags Http/Project
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/project/create [post]
func create(ctx *gin.Context) {
	var (
		input   = new(project.CreateProjectReq)
		session = gopowersso.GetSession(ctx)
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
