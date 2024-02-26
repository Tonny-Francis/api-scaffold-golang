package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response interface {
	Ok(json interface{}) HttpResponse
	Created(json interface{}) HttpResponse
	NoContent() HttpResponse
}

type Error interface {
	BadRequest(message string) HttpError
	Unauthorized(message string) HttpError
	Forbidden(message string) HttpError
	NotFound(message string) HttpError
	InternalServerError(message string) HttpError
	Duplicated(message string) HttpError
	UnprocessableEntity(message string) HttpError
	TooManyRequests(message string) HttpError
	ServiceUnavailable(message string) HttpError
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