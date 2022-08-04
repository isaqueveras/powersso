// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/isaqueveras/lingo"
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/middleware"
	"github.com/isaqueveras/power-sso/pkg/database/postgres"
	"github.com/isaqueveras/power-sso/pkg/database/redis"
	"github.com/isaqueveras/power-sso/pkg/i18n"
	"github.com/isaqueveras/power-sso/pkg/logger"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var setupLingo = lingo.New(i18n.EnglishUS, "i18n")
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

	router := gin.New()
	router.Use(
		middleware.VersionInfo(),
		middleware.SetupI18n(setupLingo),
		middleware.RequestIdentifier(),
		middleware.RecoveryWithZap(logg.ZapLogger(), true),
		middleware.GinZap(logg.ZapLogger(), *cfg),
	)

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": i18n.Value("welcome.title"), "date": time.Now()})
	})

	group := errgroup.Group{}
	group.Go(func() error {
		return endless.ListenAndServe("0.0.0.0"+cfg.Server.Port, router)
	})

	if err = group.Wait(); err != nil {
		logg.Fatal("Error while serving the application", err)
	}
}
