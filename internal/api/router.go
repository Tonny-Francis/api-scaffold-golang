package api

import (
	"context"

	v1 "github.com/Tonny-Francis/api-base-golang/internal/api/v1"
	"github.com/Tonny-Francis/api-base-golang/internal/container"
	"github.com/gin-gonic/gin"
)

func Router(ctx context.Context, deps *container.Dependencies) *gin.Engine {
	router := deps.Components.Router

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	apiV1Router := router.Group("/v1")

	v1.Routes(ctx, deps, apiV1Router)

	return router
}
