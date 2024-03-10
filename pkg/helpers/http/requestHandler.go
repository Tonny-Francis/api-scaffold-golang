package http

import (
	"errors"
	"fmt"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (r *DefaultHandler) RequestHandler(logger *logrus.Logger, controller func(c *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		completeRoute := c.Request.Method + " " + c.FullPath()

		response, err := controller(c)

		if err != nil {
			errorMiddleware(err, c, logger)
			return
		}


		if err == nil && response == nil {
			err = errors.New("Response is nil and error is nil")
			errorMiddleware(err, c, logger)
			return
		}

		if httpResponse, ok := response.(HttpResponse); ok {
			c.JSON(httpResponse.StatusCode, httpResponse.Payload)
			return
		}

		logger.Warnf("Controller on route %v did not return a HttpResponse object", completeRoute)
		c.Status(nethttp.StatusOK)
	}
}

func errorMiddleware(err error, c *gin.Context, logger *logrus.Logger) {
	var statusCode int
	var message interface{}
	fmt.Printf("Tipo do Erro: %T\n", err)

	logger.Error(err)
	switch err := err.(type) {
	case validator.ValidationErrors:
		validationErr := err
		statusCode = nethttp.StatusBadRequest
		message = map[string]interface{}{"message": validationErr.Error()}
	case HttpError:
		httpErr := err
		statusCode = httpErr.StatusCode
		message = map[string]interface{}{"message": httpErr.Message}
	default:
		errorIdentifier := uuid.New().String()
		logger.Errorf("Identifier: %s ; Error on route %s, %s", errorIdentifier, c.FullPath(), err.Error())
		statusCode = nethttp.StatusInternalServerError
		message = map[string]interface{}{"message": "Internal server error, please contact the administrator with the identifier " + errorIdentifier}
	}

	c.JSON(statusCode, message)
}
