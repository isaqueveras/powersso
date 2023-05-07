// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gopowersso "github.com/isaqueveras/go-powersso"

	"github.com/isaqueveras/power-sso/internal/application/auth"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// register godoc
// @Summary Register a user
// @Description Register a user
// @Tags Http/Auth
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/auth/register [post]
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

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}

// activation godoc
// @Summary Activate the user
// @Description Route to activate the user
// @Tags Http/Auth
// @Accept json
// @Produce json
// @Success 200 {object} utils.NoContent{}
// @Router /v1/auth/activation/{token} [post]
func activation(ctx *gin.Context) {
	token := ctx.Param("token")

	if err := auth.Activation(ctx, &token); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, utils.NoContent{})
}

// login godoc
// @Summary User login
// @Description Route to login a user account into the system
// @Tags Http/Auth
// @Accept json
// @Produce json
// @Success 200 {object} auth.SessionResponse
// @Router /v1/auth/login [post]
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

// logout godoc
// @Summary User logout
// @Description Route to logout a user account into the system
// @Tags Http/Auth
// @Accept json
// @Produce json
// @Success 204 {object} utils.NoContent{}
// @Router /v1/auth/logout [delete]
func logout(ctx *gin.Context) {
	s := gopowersso.GetSession(ctx)
	if err := auth.Logout(ctx, &s.SessionID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusNoContent, utils.NoContent{})
}

// loginSteps godoc
// @Summary Steps to login
// @Description Steps to login
// @Tags Http/Auth
// @Accept json
// @Produce json
// @Success 200 {object} auth.StepsResponse
// @Router /v1/auth/login/steps [get]
func loginSteps(ctx *gin.Context) {
	var (
		res *auth.StepsResponse
		err error
	)

	if res, err = auth.LoginSteps(ctx, utils.GetStringPointer(ctx.Query("email"))); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
