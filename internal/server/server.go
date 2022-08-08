// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package server

import (
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/isaqueveras/endless"
	"github.com/isaqueveras/lingo"
	"github.com/isaqueveras/power-sso/config"
	"github.com/isaqueveras/power-sso/internal/middleware"
	"github.com/isaqueveras/power-sso/internal/presentation/auth"
	"github.com/isaqueveras/power-sso/pkg/i18n"
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

func (s *Server) Run() error {
	var setupLingo = lingo.New(i18n.EnglishUS, "i18n")

	if s.cfg.Server.Mode == config.ModeProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(
		middleware.VersionInfo(),
		middleware.SetupI18n(setupLingo),
		middleware.RequestIdentifier(),
		middleware.RecoveryWithZap(s.logg.ZapLogger(), true),
		middleware.GinZap(s.logg.ZapLogger(), *s.cfg),
	)

	v1 := router.Group("v1")
	auth.Router(v1.Group("auth"))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": i18n.Value("welcome.title"), "date": time.Now()})
	})

	endless.DefaultReadTimeOut = s.cfg.Server.ReadTimeout * time.Second
	endless.DefaultWriteTimeOut = s.cfg.Server.WriteTimeout * time.Second
	endless.DefaultMaxHeaderBytes = http.DefaultMaxHeaderBytes

	group := errgroup.Group{}
	group.Go(func() error {
		if s.cfg.Server.SSL {
			return endless.ListenAndServeTLS("0.0.0.0"+s.cfg.Server.Port, certFile, keyFile, router)
		} else {
			return endless.ListenAndServe("0.0.0.0"+s.cfg.Server.Port, router)
		}
	})

	prefixRouter := router.Group("debug/pprof")
	prefixRouter.GET("/", func(c *gin.Context) {
		// TODO: only admin can access this
		pprof.Index(c.Writer, c.Request)
	})

	group.Go(func() error {
		return endless.ListenAndServe("0.0.0.0"+s.cfg.Server.PprofPort, router)
	})

	if err := group.Wait(); err != nil {
		s.logg.Fatal("Error while serving the application: ", err)
	}

	return nil
}
