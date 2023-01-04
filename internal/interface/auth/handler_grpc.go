// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import (
	"context"

	"github.com/isaqueveras/power-sso/internal/application/auth"
)

// Server implements proto interface
type Server struct {
	auth.UnimplementedAuthenticationServer
}

// CreateUser create user using gRPC
func (s *Server) registerUser(ctx context.Context, in *auth.User) (*auth.Empty, error) {
	return nil, nil
}
