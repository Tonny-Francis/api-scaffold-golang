package v1

import (
	"context"

	"github.com/Tonny-Francis/api-base-golang/internal/container"
	"github.com/gin-gonic/gin"
)

func Routes(ctx context.Context, deps *container.Dependencies, apiV1Router *gin.RouterGroup) {
	requestHandler := deps.Services.RequestHandler.RequestHandler
	logger := deps.Components.Logger
	// Router health check
	apiV1Router.GET("/example", requestHandler(logger, func(c *gin.Context) (interface{}, error) {
		return gin.H{"status": "ok"}, nil
	}))
}
