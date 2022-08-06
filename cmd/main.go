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

	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration file: ", err)
	}

	var cfg *config.Config
	if cfg, err = config.ParseConfig(cfgFile); err != nil {
		log.Fatal("Error parsing configuration file: ", err)
	}

	logg := logger.NewLogger(cfg)
	logg.InitLogger()

	if err = postgres.OpenConnections(cfg); err != nil {
		logg.Fatal("Unable to open connections to database: ", err)
	}
	defer postgres.CloseConnections()

	var redisClient = redis.NewRedisClient(cfg)
	defer redisClient.Close()

	if server.NewServer(cfg, logg).Run(); err != nil {
		logg.Fatal("Error while serving the application: ", err)
	}
}
