package requestHandlerHelper

import (
	"fmt"
	"net/http"

	"github.com/Tonny-Francis/api-base-golang/helpers/httpErrorHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/httpResponseHelper"
	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func RequestHandler(controller func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		completeRoute := c.Request.Method + " " + c.FullPath()

		response, err := controller(c)

		if err != nil {
			errorMiddleware(err, c)
			return
		}

		if httpResponse, ok := response.(httpResponseHelper.HttpResponse); ok {
			c.JSON(httpResponse.StatusCode, httpResponse.Payload)
			return
		}

		loggerHelper.Logger.Warnf("Controller on route %v did not return a HttpResponse object", completeRoute)
		c.Status(http.StatusOK)
	}
}

func errorMiddleware(err error, c *gin.Context) {
	var statusCode int
	var message interface{}
	fmt.Printf("Tipo do Erro: %T\n", err)

	loggerHelper.Logger.Error(err)
	switch err.(type) {
	case validator.ValidationErrors:
		validationErr := err.(validator.ValidationErrors)
		statusCode = http.StatusBadRequest
		message = map[string]interface{}{"message": validationErr.Error()}
	case httpErrorHelper.HttpError:
		httpErr := err.(httpErrorHelper.HttpError)
		statusCode = httpErr.StatusCode
		message = map[string]interface{}{"message": httpErr.Message}
	default:
		errorIdentifier := uuid.New().String()
		loggerHelper.Logger.Errorf("Identifier: %s ; Error on route %s, %s", errorIdentifier, c.FullPath(), err.Error())
		statusCode = http.StatusInternalServerError
		message = map[string]interface{}{"message": "Internal server error, please contact the administrator with the identifier " + errorIdentifier}
	}

	c.JSON(statusCode, message)
}
