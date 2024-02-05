package httpResponseHelper

import (
	"github.com/gin-gonic/gin"
)

var ResponseCodes = map[string]int{
	"OK": 200,
	"CREATED": 201,
	"NO_CONTENT": 204,
}

type HttpResponse struct {
	StatusCode int
	Payload    interface{}
}

func httpResponse(c *gin.Context, response HttpResponse) {
	c.JSON(response.StatusCode, response.Payload)
}

func Ok(c *gin.Context, json interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["OK"],
		Payload: json,
	}
}

func Created(c *gin.Context, json interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["CREATED"],
		Payload: json,
	}
}

func NoContent(c *gin.Context) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["NO_CONTENT"],
		Payload: nil,
	}
}