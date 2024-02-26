package http

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

func (r *DefaultResponse) Ok(c *gin.Context, json interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["OK"],
		Payload: json,
	}
}

func (r *DefaultResponse) Created(c *gin.Context, json interface{}) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["CREATED"],
		Payload: json,
	}
}

func (r *DefaultResponse) NoContent(c *gin.Context) HttpResponse {
	return HttpResponse{
		StatusCode: ResponseCodes["NO_CONTENT"],
		Payload: nil,
	}
}