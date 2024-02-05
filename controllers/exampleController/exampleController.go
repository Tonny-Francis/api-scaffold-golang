package exampleController

import (
	"github.com/Tonny-Francis/api-base-golang/helpers/httpResponseHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) (interface{}, error) {
	loggerHelper.Logger.Info("Hello World")

	return httpResponseHelper.Ok(c, gin.H{"message": "Hello World"}), nil
}
