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

	// Router health check
	apiV1Router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Routers

	exampleRouter.ExampleRouter(apiV1Router)

	// Start server

	err := router.Run(":" + envConfig.GetEnvConfig().PORT)

	if err != nil {
		loggerHelper.Logger.Errorf("Error initializing server: %v", err)
	}
}
