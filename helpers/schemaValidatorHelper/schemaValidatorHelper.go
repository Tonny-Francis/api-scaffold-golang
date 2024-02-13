package schemaValidatorHelper

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Validator faz o bind do corpo da requisição à struct e a valida
func Validator(c *gin.Context, data interface{}) (interface{}, error) {
	// BindJSON para analisar e vincular o corpo da requisição à struct
	if err := c.BindJSON(data); err != nil {
		return nil, err
	}

	// Realize a validação da struct
	if err := validate.Struct(data); err != nil {
		return nil, err
	}
	
	// Retorna o corpo validado
	return data, nil
}
