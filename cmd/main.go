package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	var router = gin.New()

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to PowerSSO",
			"date":    time.Now(),
		})
	})

	group := errgroup.Group{}
	group.Go(func() error {
		return endless.ListenAndServe("0.0.0.0:5820", router)
	})

	if err := group.Wait(); err != nil {
		log.Fatal("Error while serving the application", err.Error())
	}
}
