package example

import (
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/http"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (r *DefaultHandler) Get(httResponse http.Response, httpError http.Error, logger *logrus.Logger, parseSchema validator.SchemaValidator) func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var body MyForm
		var header MyHeaders
		var uri MyParams
		var query MyQueryParams

		err := parseSchema.SchemaValidator(c, map[string]interface{}{
			"body":   &body,
			"uri":    &uri,
			"header": &header,
			"query":  &query,
		})

		if err != nil {
			return nil, err
		}

		return httResponse.Ok(gin.H{"message": gin.H{
			"URI":    uri.ID,
			"Name":   body.Name,
			"Phone":  body.Phone,
			"Tenant": header.Tenant,
			"Page":   query.Page,
		}}), nil
	}
}
