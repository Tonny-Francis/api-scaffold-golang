package exampleController

import (
	"github.com/Tonny-Francis/api-base-golang/helpers/httpResponseHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/schemaValidatorHelper"
	"github.com/Tonny-Francis/api-base-golang/schemas/exampleSchema"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) (interface{}, error) {
	body, err := schemaValidatorHelper.Validator(c, &exampleSchema.ExampleSchema{})
	
	parseBody := body.(*exampleSchema.ExampleSchema)

	if err != nil {
		return nil, err
	}

	// Se a validação foi bem-sucedida, continue com o processamento
	return httpResponseHelper.Ok(c, gin.H{
		"message": "Hello World",
		"body": parseBody,
	}), nil
}
