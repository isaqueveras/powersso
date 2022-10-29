// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gopowersso "github.com/isaqueveras/go-powersso"

	"github.com/isaqueveras/power-sso/internal/application/project"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

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

	ctx.JSON(http.StatusCreated, nil)
}
