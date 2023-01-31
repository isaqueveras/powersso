// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/isaqueveras/power-sso/internal/application/auth"
	"github.com/isaqueveras/power-sso/internal/utils"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

// Server implements proto interface
type Server struct {
	UnimplementedAuthenticationServer
}

// RegisterUser register user
func (s *Server) RegisterUser(ctx context.Context, in *User) (_ *Empty, err error) {
	postcode, err := strconv.Atoi(in.GetPostCode())
	if err != nil {
		return nil, oops.HandlingGRPC(err)
	}

	birthday, err := time.Parse(time.RFC3339, in.Birthday)
	if err != nil {
		return nil, oops.HandlingGRPC(err)
	}

	if err = auth.Register(ctx, &auth.RegisterRequest{
		FirstName:   utils.GetStringPointer(in.FirstName),
		LastName:    utils.GetStringPointer(in.LastName),
		Email:       utils.GetStringPointer(in.Email),
		Password:    utils.GetStringPointer(in.Password),
		About:       utils.GetStringPointer(in.About),
		Avatar:      utils.GetStringPointer(in.Avatar),
		PhoneNumber: utils.GetStringPointer(in.PhoneNumber),
		Address:     utils.GetStringPointer(in.Address),
		City:        utils.GetStringPointer(in.City),
		Country:     utils.GetStringPointer(in.Country),
		Gender:      utils.GetStringPointer(in.Gender),
		Postcode:    utils.GetIntPointer(postcode),
		Birthday:    utils.GetTimePointer(birthday),
	}); err != nil {
		return nil, oops.HandlingGRPC(err)
	}

	return
}
