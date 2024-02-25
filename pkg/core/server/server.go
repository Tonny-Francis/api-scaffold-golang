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

func Run(ctx context.Context, deps *container.Dependencies, router *gin.Engine) {
	server := &http.Server{
		Addr:    ":" + deps.Components.Env.PORT,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			deps.Components.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	deps.Components.Logger.Infof("Server started on port %s\n", deps.Components.Env.PORT)

	sig := make(chan os.Signal)

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-sig

	deps.Components.Logger.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		deps.Components.Logger.Fatalf("Server forced to shutdown: %s\n", err)
	}

	select {
	case <-ctx.Done():
		deps.Components.Logger.Println("Server timeout")
	}

	deps.Components.Logger.Println("Server exiting")
}