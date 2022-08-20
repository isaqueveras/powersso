// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"

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

	ctx.JSON(http.StatusCreated, gin.H{})
}

func activation(ctx *gin.Context) {
	token := ctx.Param("token")

	if err := auth.Activation(ctx, &token); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}

func login(ctx *gin.Context) {
	var (
		input  auth.LoginRequest
		output *auth.SessionResponse
		err    error
	)

	if err = ctx.ShouldBindJSON(&input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	input.ClientIP = ctx.ClientIP()
	input.UserAgent = ctx.Request.UserAgent()
	input.Validate()

	if output, err = auth.Login(ctx, &input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}
