package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouters() {
	router := gin.Default()

	apiV1Router := router.Group("/v1")

	exampleRouter(apiV1Router)

	router.Run(":8000")
}