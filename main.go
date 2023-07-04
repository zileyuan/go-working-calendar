package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zileyuan/go-working-calendar/config"
	"github.com/zileyuan/go-working-calendar/middleware"
	"github.com/zileyuan/go-working-calendar/router"

	"github.com/zileyuan/go-working-calendar/util"
)

func startHTTP() {
	if strings.ToLower(config.Runtime.Common.RunMode) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(gin.LoggerWithWriter(util.FileWriter))
	g.Use(gin.RecoveryWithWriter(util.FileWriter))
	g.Use(middleware.CrossAccessAllow())
	g.Use(middleware.AuthToken())
	router.Route(g)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", config.Runtime.RunEnv.HTTPPort),
		Handler: g,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			util.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	util.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		util.Fatalf("Server Shutdown: %s\n", err)
	}
	util.Info("Server exiting")
}

func main() {
	//create http instance
	startHTTP()
}
