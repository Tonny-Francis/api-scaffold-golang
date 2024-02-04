package server

import (
	"github.com/Tonny-Francis/api-base-golang/configs/envConfig"
	"github.com/Tonny-Francis/api-base-golang/routers/exampleRouter"
	"github.com/gin-gonic/gin"
	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
)

func InitRouters() {
	// Initialize Gin Router and set mode
	gin.SetMode(envConfig.GetEnvConfig().GIN_MODE)

	router := gin.New()

	apiV1Router := router.Group("/v1")

	exampleRouter.ExampleRouter(apiV1Router)

	err := router.Run(":" + envConfig.GetEnvConfig().PORT)

	if err != nil {
		loggerHelper.Logger.Errorf("error initializing router: %v", err)
	}
}
