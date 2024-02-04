package server

import (
	"github.com/Tonny-Francis/api-base-golang/routers/exampleRouter"
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	router := gin.Default()

	apiV1Router := router.Group("/v1")

	exampleRouter.ExampleRouter(apiV1Router)

	router.Run(":8000")
}
