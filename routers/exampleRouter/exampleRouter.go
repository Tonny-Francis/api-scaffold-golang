package exampleRouter

import (
	"github.com/Tonny-Francis/api-base-golang/controllers/exampleController"
	"github.com/Tonny-Francis/api-base-golang/helpers/requestHandlerHelper"

	//"github.com/Tonny-Francis/api-base-golang/middlewares/authMiddleware"
	"github.com/gin-gonic/gin"
)

func ExampleRouter(RouterGroup *gin.RouterGroup) {
	//Use Middleware

	// RouterGroup.Use(
	// 	authMiddleware.Auth(),
	// )

	RouterGroup.POST("/example", requestHandlerHelper.RequestHandler(exampleController.Get))
}
