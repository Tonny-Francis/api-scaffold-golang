package handler

import (
	"github.com/Tonny-Francis/api-base-golang/helper/httpResponse"
	"github.com/gin-gonic/gin"
)

func ExampleHandler(c *gin.Context) {
	httpResponse.Ok(c, gin.H{"message": "Hello World"})
}
