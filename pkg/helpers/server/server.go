package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Tonny-Francis/api-base-golang/internal/container"
	"github.com/gin-gonic/gin"
)

func (r *DefaultServer) Run(ctx context.Context, helpers *container.Helpers, services *container.Services, domains *container.Domains, router *gin.Engine) {
	server := &http.Server{
		Addr:    ":" + helpers.Env.PORT,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			helpers.Logger.Fatalf("Listen: %s\n", err)
		}
	}()

	helpers.Logger.Infof("Server started on port %s\n", helpers.Env.PORT)

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	helpers.Logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		helpers.Logger.Fatalf("Server forced to shutdown: %s\n", err)
	}

	<-sig

	helpers.Logger.Println("Server exiting")
}