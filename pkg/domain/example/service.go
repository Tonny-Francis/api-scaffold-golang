package example

import (
	"github.com/Tonny-Francis/api-base-golang/pkg/services/http"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Get(httpResponse *http.DefaultResponse, c *gin.Context) (interface{}, error)
}

type DefaultHandler struct{}

func NewHandler() Handler {
	return &DefaultHandler{}
}
