package exampleController

import (
	"github.com/Tonny-Francis/api-base-golang/helpers/httpErrorHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) (interface{}, error) {
	loggerHelper.Logger.Info("Hello World")

	return nil, httpErrorHelper.BadRequest(c, "Hello World")
}
