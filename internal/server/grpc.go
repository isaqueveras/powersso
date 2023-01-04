// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package server

import (
	"net"

	gogrpc "google.golang.org/grpc"

	authApp "github.com/isaqueveras/power-sso/internal/application/auth"
	authInterface "github.com/isaqueveras/power-sso/internal/interface/auth"
	"github.com/isaqueveras/power-sso/pkg/oops"
)

func (s *Server) RunGRPC() (err error) {
	var (
		listen net.Listener
		server = gogrpc.NewServer()
	)

	if listen, err = net.Listen("tcp", "localhost:50050"); err != nil {
		return oops.Err(err)
	}

	authApp.RegisterAuthenticationServer(server, &authInterface.Server{})

	if err = server.Serve(listen); err != nil {
		return oops.Err(err)
	}

	return
}
