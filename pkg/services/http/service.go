package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response interface {
	Ok(c *gin.Context, json interface{}) HttpResponse
	Created(c *gin.Context, json interface{}) HttpResponse
	NoContent(c *gin.Context) HttpResponse
}

type Error interface {
	BadRequest(c *gin.Context, message string) HttpError
	Unauthorized(c *gin.Context, message string) HttpError
	Forbidden(c *gin.Context, message string) HttpError
	NotFound(c *gin.Context, message string) HttpError
	InternalServerError(c *gin.Context, message string) HttpError
	Duplicated(c *gin.Context, message string) HttpError
	UnprocessableEntity(c *gin.Context, message string) HttpError
	TooManyRequests(c *gin.Context, message string) HttpError
	ServiceUnavailable(c *gin.Context, message string) HttpError
}

type Handler interface {
	RequestHandler(logger *logrus.Logger, controller func(c *gin.Context) (interface{}, error)) gin.HandlerFunc
}

type DefaultError struct{}

type DefaultResponse struct{}

type DefaultHandler struct{}

func NewResponseService() Response {
	return &DefaultResponse{}
}

func NewErrorService() Error {
	return &DefaultError{}
}

func NewRequestHandler() Handler {
	return &DefaultHandler{}
}