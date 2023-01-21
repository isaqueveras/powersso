// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package auth

import "context"

// Server implements proto interface
type Server struct {
	UnimplementedAuthenticationServer
}

// CreateUser create user using gRPC
func (s *Server) registerUser(ctx context.Context, in *User) (*Empty, error) {
	return nil, nil
}
