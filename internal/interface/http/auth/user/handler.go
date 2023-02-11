// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/isaqueveras/power-sso/internal/application/auth/user"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

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
	userUUID, err := uuid.Parse(ctx.Param("user_uuid"))
	if err != nil {
		oops.Handling(ctx, err)
		return
	}

	if err = user.Disable(ctx, &userUUID); err != nil {
		oops.Handling(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, utils.NoContent{})
}
