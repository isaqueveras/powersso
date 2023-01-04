// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"

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

	server := server.NewServer(cfg, logg)

	if err = server.RunRest(); err != nil {
		logg.Fatal("Error while serving the application: ", err)
	}

	if err = server.RunGRPC(); err != nil {
		logg.Fatal("Error while serving the server gRPC: ", err)
	}
}
