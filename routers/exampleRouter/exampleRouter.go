package exampleRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/Tonny-Francis/api-base-golang/controllers/exampleController"
)

func ExampleRouter(RouterGroup *gin.RouterGroup){
	RouterGroup.GET("/example", exampleController.Get)
}