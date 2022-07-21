package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/middleware"
	"github.com/isaqueveras/power-sso/pkg/logger"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		logg    *zap.Logger
		cfgFile *viper.Viper
		cfg     *config.Config
		err     error
	)

	if cfgFile, err = config.LoadConfig(); err != nil {
		log.Fatal("Error loading configuration file: ", err)
	}

	if cfg, err = config.ParseConfig(cfgFile); err != nil {
		log.Fatal("Error parsing configuration file: ", err)
	}

	if logg, err = logger.NewLogger(cfg).InitLogger(); err != nil {
		log.Fatal(err)
	}
	defer func() { _ = logg.Sync() }()
	zap.ReplaceGlobals(logg)

	var router = gin.New()
	router.Use(
		ginzap.RecoveryWithZap(logg, true),
		middleware.GinZap(logg, *cfg),
	)

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to PowerSSO", "date": time.Now()})
	})

	group := errgroup.Group{}
	group.Go(func() error {
		return endless.ListenAndServe("0.0.0.0"+cfg.Server.Port, router)
	})

	if err := group.Wait(); err != nil {
		logg.Fatal("Error while serving the application", zap.Error(err))
	}
}
