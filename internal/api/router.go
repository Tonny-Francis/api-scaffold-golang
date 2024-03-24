package api

import (
	"context"

	v1 "github.com/Tonny-Francis/api-scaffold-golang/internal/api/v1"
	"github.com/Tonny-Francis/api-scaffold-golang/internal/container"
	"github.com/gin-gonic/gin"
)

func Router(ctx context.Context, helpers *container.Helpers, services *container.Services, domains *container.Domains) *gin.Engine {
	router := helpers.Router

	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	apiV1Router := router.Group("/v1")

	v1.Routes(ctx, helpers, services, domains, apiV1Router)

	return router
}
