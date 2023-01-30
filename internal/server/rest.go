// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT style
// license that can be found in the LICENSE file.

package server

import (
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaqueveras/endless"
	gopowersso "github.com/isaqueveras/go-powersso"

	"github.com/isaqueveras/power-sso/internal/interface/http/auth"
	"github.com/isaqueveras/power-sso/internal/interface/project"
	"github.com/isaqueveras/power-sso/internal/middleware"
	"github.com/isaqueveras/power-sso/pkg/i18n"
)

func (s *Server) ServerHTTP() (err error) {
	if !s.cfg.Server.StartHTTP {
		return
	}

	s.logg.Info("Server HTTP is running")
	if s.cfg.Server.IsModeProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(
		middleware.CORS(),
		middleware.VersionInfo(),
		middleware.SetupI18n(),
		middleware.RequestIdentifier(),
		middleware.RecoveryWithZap(s.logg.ZapLogger(), true),
		middleware.GinZap(s.logg.ZapLogger(), *s.cfg),
	)

	v1 := router.Group("v1")
	auth.Router(v1.Group("auth"))
	project.RouterAuthorization(v1.Group("project", gopowersso.Authorization(&s.cfg.UserAuthToken.SecretKey)))
	auth.RouterAuthorization(v1.Group("auth", gopowersso.Authorization(&s.cfg.UserAuthToken.SecretKey)))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": i18n.Value("welcome.title"), "date": time.Now()})
	})

	endless.DefaultReadTimeOut = s.cfg.Server.ReadTimeout * time.Second
	endless.DefaultWriteTimeOut = s.cfg.Server.WriteTimeout * time.Second
	endless.DefaultMaxHeaderBytes = http.DefaultMaxHeaderBytes

	s.group.Go(func() error {
		if s.cfg.Server.SSL {
			return endless.ListenAndServeTLS("0.0.0.0"+s.cfg.Server.Port, certFile, keyFile, router)
		} else {
			return endless.ListenAndServe("0.0.0.0"+s.cfg.Server.Port, router)
		}
	})

	s.routerDebugPProf(router)

	return nil
}

func (s *Server) routerDebugPProf(router *gin.Engine) {
	prefixRouter := router.Group("debug/pprof")
	prefixRouter.GET("/",
		gopowersso.Authorization(&s.cfg.UserAuthToken.SecretKey),
		gopowersso.OnlyAdmin(),
		func(c *gin.Context) {
			pprof.Index(c.Writer, c.Request)
		},
	)

	s.group.Go(func() error {
		return endless.ListenAndServe("0.0.0.0"+s.cfg.Server.PprofPort, router)
	})
}
