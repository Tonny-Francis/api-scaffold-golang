package v1

import (
	"context"

	"github.com/Tonny-Francis/api-scaffold-golang/internal/container"
	"github.com/gin-gonic/gin"
)

func Routes(ctx context.Context, helpers *container.Helpers, services *container.Services, domains *container.Domains, apiV1Router *gin.RouterGroup) {
	apiV1Router.POST("/example/:id", helpers.RequestHandler.RequestHandler(helpers.Logger, domains.Example.Get(helpers.HttpResponse, helpers.HttpError, helpers.Logger, helpers.ParseSchema)))
}
