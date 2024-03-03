package v1

import (
	"context"

	"github.com/Tonny-Francis/api-base-golang/internal/container"
	"github.com/gin-gonic/gin"
)

func Routes(ctx context.Context, deps *container.Dependencies, apiV1Router *gin.RouterGroup) {
	// Dependencies
	requestHandler := deps.Services.RequestHandler.RequestHandler
	logger := deps.Components.Logger
	example := deps.Services.Example
	httpResponse := deps.Services.HttpResponse
	httpError := deps.Services.HttpError
	parseSchema := deps.Services.ParseSchema

	// Router health check
	apiV1Router.POST("/example/:id", requestHandler(logger, example.Get(httpResponse, httpError, logger, parseSchema)))
}
