// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/pkg/logger"
)

const (
	certFile = "ssl/server.crt"
	keyFile  = "ssl/server.pem"
)

// Server struct
type Server struct {
	cfg  *config.Config
	logg *logger.Logger
}

// NewServer new server constructor
func NewServer(cfg *config.Config, logg *logger.Logger) *Server {
	return &Server{
		cfg:  cfg,
		logg: logg,
	}
}
