package example

import (
	"github.com/Tonny-Francis/api-base-golang/pkg/services/http"
	"github.com/Tonny-Francis/api-base-golang/pkg/services/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	Get(httpResponse http.Response, httpError http.Error, logger *logrus.Logger, parseSchema validator.SchemaValidator) func(c *gin.Context) (interface{}, error)
}

type DefaultHandler struct{}

func NewHandler() Handler {
	return &DefaultHandler{}
}
