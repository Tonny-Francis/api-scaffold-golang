package exampleRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/Tonny-Francis/api-base-golang/controllers/exampleController"
	"github.com/Tonny-Francis/api-base-golang/helpers/requestHandlerHelper"
)

func ExampleRouter(RouterGroup *gin.RouterGroup){
	RouterGroup.GET("/example", requestHandlerHelper.RequestHandler(exampleController.Get))
}