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

	"github.com/isaqueveras/powersso/delivery/http/auth"
	"github.com/isaqueveras/powersso/delivery/http/permissions"
	"github.com/isaqueveras/powersso/delivery/http/project"
	"github.com/isaqueveras/powersso/middleware"
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
		middleware.RequestIdentifier(),
		middleware.RecoveryWithZap(s.logg.ZapLogger(), true),
		middleware.GinZap(s.logg.ZapLogger(), *s.cfg),
		middleware.SetupI18n(),
	)

	v1 := router.Group("v1")
	auth.Router(v1.Group("auth"))
	auth.RouterAuthorization(v1.Group("auth", middleware.Auth()))
	project.Router(v1.Group("project", middleware.Auth()))
	permissions.Router(v1.Group("permission", middleware.Auth()))

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

	if !s.cfg.Server.IsModeDevelopment() {
		s.routerDebugPProf(router)
	}

	return
}

func (s *Server) routerDebugPProf(router *gin.Engine) {
	r := router.Group("debug/pprof")
	r.Use(middleware.Auth(), middleware.OnlyAdmin())
	r.GET("/", func(c *gin.Context) { pprof.Index(c.Writer, c.Request) })
	r.GET("/cmdline", func(c *gin.Context) { pprof.Cmdline(c.Writer, c.Request) })
	r.GET("/profile", func(c *gin.Context) { pprof.Profile(c.Writer, c.Request) })
	r.POST("/symbol", func(c *gin.Context) { pprof.Symbol(c.Writer, c.Request) })
	r.GET("/trace", func(c *gin.Context) { pprof.Trace(c.Writer, c.Request) })
	r.GET("/allocs", func(c *gin.Context) { pprof.Handler("allocs").ServeHTTP(c.Writer, c.Request) })
	r.GET("/block", func(c *gin.Context) { pprof.Handler("block").ServeHTTP(c.Writer, c.Request) })
	r.GET("/goroutine", func(c *gin.Context) { pprof.Handler("goroutine").ServeHTTP(c.Writer, c.Request) })
	r.GET("/heap", func(c *gin.Context) { pprof.Handler("heap").ServeHTTP(c.Writer, c.Request) })
	r.GET("/mutex", func(c *gin.Context) { pprof.Handler("mutex").ServeHTTP(c.Writer, c.Request) })
	r.GET("/threadcreate", func(c *gin.Context) { pprof.Handler("threadcreate").ServeHTTP(c.Writer, c.Request) })
	s.group.Go(func() error { return endless.ListenAndServe("0.0.0.0"+s.cfg.Server.PprofPort, router) })
}
