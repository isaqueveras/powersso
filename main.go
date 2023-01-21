// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/server"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/database/redis"
	"github.com/isaqueveras/power-sso/pkg/logger"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.LoadConfig(".")

	var (
		cfg  = config.Get()
		logg = logger.NewLogger(cfg)
		err  error
	)

	logg.InitLogger()
	if err = postgres.OpenConnections(cfg); err != nil {
		logg.Fatal("Unable to open connections to database: ", err)
	}
	defer postgres.CloseConnections()

	redis := redis.NewRedisClient(cfg)
	defer redis.Close()

	var (
		group  = &errgroup.Group{}
		server = server.NewServer(cfg, logg, group)
	)

	// TODO: add in the configuration if it is to run http server
	if err = server.ServerHTTP(); err != nil {
		logg.Fatal("Error while serving the server HTTP: ", err)
	}

	// TODO: add in the configuration if it is to run grpc server
	if err = server.ServerGRPC(); err != nil {
		logg.Fatal("Error while serving the server GRPC: ", err)
	}

	if err = group.Wait(); err != nil {
		logg.Fatal("Error while serving the servers: ", err)
	}
}
