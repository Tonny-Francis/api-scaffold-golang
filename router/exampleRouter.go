package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Tonny-Francis/api-base-golang/handler"
)

func exampleRouter(RouterGroup *gin.RouterGroup){
	RouterGroup.GET("/example", handler.ExampleHandler)
}