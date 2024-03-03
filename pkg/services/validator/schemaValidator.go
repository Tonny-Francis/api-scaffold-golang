package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func (r *DefaultSchemaValidator) SchemaValidator(c *gin.Context, bindSchemas map[string]interface{}) error {
	for bindType, schema := range bindSchemas {
		switch bindType {
		case "body":
			if err := c.ShouldBind(schema); err != nil {
				return err
			}
		case "uri":
			if err := c.ShouldBindUri(schema); err != nil {
				return err
			}
		case "header":
			if err := c.ShouldBindHeader(schema); err != nil {
				return err
			}
		case "query":
			if err := c.ShouldBindQuery(schema); err != nil {
				return err
			}
		default:
			err := errors.New("unsupported bind type")

			return err
		}
	}

	return nil
}
