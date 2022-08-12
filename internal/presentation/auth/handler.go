// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/isaqueveras/power-sso/internal/application/auth"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

func register(ctx *gin.Context) {
	var (
		input auth.RegisterRequest
		err   error
	)

	if err = ctx.ShouldBindJSON(&input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err = auth.Register(ctx, &input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(201, nil)
}

func activation(ctx *gin.Context) {
	token := ctx.Param("token")

	if err := auth.Activation(ctx, &token); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(201, gin.H{})
}
