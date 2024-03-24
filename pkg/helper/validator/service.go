package validator

import "github.com/gin-gonic/gin"

type SchemaValidator interface {
	SchemaValidator(c *gin.Context,  bindSchemas map[string]interface{}) error
}

type DefaultSchemaValidator struct{}

func NewSchemaValidatorService() SchemaValidator {
	return &DefaultSchemaValidator{}
}
