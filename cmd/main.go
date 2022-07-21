package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/power-sso/config"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		cfgFile *viper.Viper
		cfg     *config.Config
		err     error
	)

	if cfgFile, err = config.LoadConfig(); err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	if cfg, err = config.ParseConfig(cfgFile); err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	var router = gin.New()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to PowerSSO", "date": time.Now()})
	})

	group := errgroup.Group{}
	group.Go(func() error {
		return endless.ListenAndServe("0.0.0.0"+cfg.Server.Port, router)
	})

	if err := group.Wait(); err != nil {
		log.Fatal("Error while serving the application", err.Error())
	}
}
