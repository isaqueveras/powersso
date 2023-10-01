// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/isaqueveras/powersso/config"
	"github.com/isaqueveras/powersso/database/postgres"
	_ "github.com/isaqueveras/powersso/docs"
	"github.com/isaqueveras/powersso/scripts"
	"github.com/isaqueveras/powersso/server"
	"github.com/isaqueveras/powersso/utils"
)

// @title Documentation PowerSSO
// @version 1.0.0
// @description This is PowerSSO app server.

// @contact.name PowerSSO
// @contact.url https://github.com/isaqueveras/powersso/issues

// @license.name MIT license
// @license.url https://github.com/isaqueveras/powersso/blob/main/LICENSE

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:5000
// @BasePath /
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config.LoadConfig()
	cfg := config.Get()

	logg := utils.NewLogger(cfg)
	logg.InitLogger()

	postgres.OpenConnections(cfg)
	defer postgres.CloseConnections()

	scripts.Init(logg)

	group := &errgroup.Group{}
	server := server.NewServer(cfg, logg, group)

	if err := server.ServerHTTP(); err != nil {
		logg.Fatal("Error while serving the server HTTP: ", err)
	}

	if err := server.ServerGRPC(); err != nil {
		logg.Fatal("Error while serving the server GRPC: ", err)
	}

	if err := group.Wait(); err != nil {
		logg.Fatal("Error while serving the servers: ", err)
	}
}
