// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	gopowersso "github.com/isaqueveras/go-powersso"

	app "github.com/isaqueveras/powersso/application/auth"
	domain "github.com/isaqueveras/powersso/domain/auth"
	"github.com/isaqueveras/powersso/oops"
	"github.com/isaqueveras/powersso/utils"
)

// register godoc
// @Summary Register a user
// @Description Register a user
// @Tags Http/Auth
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/auth/create_account [post]
func createAccount(ctx *gin.Context) {
	var input domain.CreateAccount
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err := app.CreateAccount(ctx, &input); err != nil {
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
	token, err := uuid.Parse(ctx.Param("token"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err := app.Activation(ctx, utils.Pointer(token)); err != nil {
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
// @Success 200 {object} auth.Session
// @Router /v1/auth/login [post]
func login(ctx *gin.Context) {
	var input domain.Login
	if err := ctx.ShouldBindJSON(&input); err != nil {
		oops.Handling(ctx, err)
		return
	}

	input.ClientIP = utils.Pointer(ctx.ClientIP())
	input.UserAgent = utils.Pointer(ctx.Request.UserAgent())
	input.Validate()

	output, err := app.Login(ctx, &input)
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, output)
}

func resetPassword(ctx *gin.Context) {
	in := &domain.ChangePassword{}
	if err := ctx.ShouldBindJSON(in); err != nil {
		oops.Handling(ctx, err)
		return
	}

	if ok := in.ValidatePassword(); !ok {
		oops.Handling(ctx, oops.New("Invalid passwords"))
		return
	}

	if err := app.ChangePassword(ctx, in); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
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
	sessionID, err := uuid.Parse(gopowersso.GetSession(ctx).SessionID)
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err := app.Logout(ctx, utils.Pointer(sessionID)); err != nil {
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
// @Success 200 {object} auth.Steps
// @Router /v1/auth/login/steps [get]
func loginSteps(ctx *gin.Context) {
	res, err := app.LoginSteps(ctx, utils.Pointer(ctx.Query("email")))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// disable godoc
// @Summary Disable user
// @Description Route to disable a user
// @Tags Http/Auth/User
// @Param user_uuid path string true "UUID of the user"
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/auth/user/{user_uuid}/disable [put]
func disable(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_uuid"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err = app.Disable(ctx, &userID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}

// configure godoc
// @Summary Configure a user's OTP
// @Description Configure a user's OTP
// @Tags Http/Auth/OTP
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/auth/user/{user_uuid}/otp/configure [post]
func configure(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_uuid"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err = app.Configure(ctx, &userID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}

// unconfigure godoc
// @Summary unconfigure a user's OTP
// @Description unconfigure a user's OTP
// @Tags Http/Auth/OTP
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/auth/user/{user_uuid}/otp/unconfigure [put]
func unconfigure(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_uuid"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err = app.Unconfigure(ctx, &userID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}

// qrcode godoc
// @Summary Configure a user's OTP
// @Description Configure a user's OTP
// @Tags Http/Auth/OTP
// @Accept json
// @Produce json
// @Success 200 {object} auth.QRCode
// @Router /v1/auth/user/{user_uuid}/otp/qrcode [get]
func qrcode(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("user_uuid"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	var res *domain.QRCode
	if res, err = app.GetQrCode(ctx, &userID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}