// Copyright (c) 2023 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package server

import (
	"golang.org/x/sync/errgroup"

	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/utils"
)

const (
	certFile = "ssl/server.crt"
	keyFile  = "ssl/server.pem"
)

// Server struct
type Server struct {
	cfg   *config.Config
	logg  *utils.Logger
	group *errgroup.Group
}

// NewServer new server constructor
func NewServer(cfg *config.Config, logg *utils.Logger, group *errgroup.Group) *Server {
	return &Server{cfg: cfg, logg: logg, group: group}
}
