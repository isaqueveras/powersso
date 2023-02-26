// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package otp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/isaqueveras/power-sso/internal/application/auth/user/otp"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// configure godoc
// @Summary Configure a user's OTP
// @Description Configure a user's OTP
// @Tags Http/Auth/OTP
// @Accept json
// @Produce json
// @Success 201 {object} utils.NoContent{}
// @Router /v1/auth/user/{user_uuid}/otp/configure [post]
func configure(ctx *gin.Context) {
	var err error
	if err = otp.Configure(ctx); err != nil {
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
// @Success 200 {object} otp.QRCodeResponse
// @Router /v1/auth/user/{user_uuid}/otp/qrcode [get]
func qrcode(ctx *gin.Context) {
	var (
		err    error
		userID uuid.UUID
		res    *otp.QRCodeResponse
	)

	if userID, err = uuid.Parse(ctx.Param("user_uuid")); err != nil {
		oops.Handling(ctx, err)
		return
	}

	if res, err = otp.GetQRCode(ctx, &userID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
