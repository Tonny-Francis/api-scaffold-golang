package example

import (
	"github.com/Tonny-Francis/api-base-golang/pkg/services/http"
	"github.com/gin-gonic/gin"
)

func (r *DefaultHandler) Get(httResponse *http.DefaultResponse, c *gin.Context) (interface{}, error) {
	return httResponse.Ok(c, gin.H{"message": "Hello, World!"}), nil
}
