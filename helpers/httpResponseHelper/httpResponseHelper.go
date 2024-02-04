package httpResponseHelper

import (
	"github.com/gin-gonic/gin"
)

var responseCodes = map[string]int{
	"OK": 200,
	"CREATED": 201,
	"NO_CONTENT": 204,
}

func httpResponse(c *gin.Context, code int, json interface{}) {
	c.JSON(code, json)
}

func Ok(c *gin.Context, json interface{}) {
	httpResponse(c, responseCodes["OK"], json)
}

func Created(c *gin.Context, json interface{}) {
	httpResponse(c, responseCodes["CREATED"], json)
}

func NoContent(c *gin.Context) {
	httpResponse(c, responseCodes["NO_CONTENT"], nil)
}