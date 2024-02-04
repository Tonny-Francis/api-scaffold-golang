package exampleController

import (
	"github.com/Tonny-Francis/api-base-golang/helpers/httpResponseHelper"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	httpResponseHelper.Ok(c, gin.H{"message": "Hello World"})
}
