package httpError

import (
	"github.com/gin-gonic/gin"
)

var errorCodes = map[string]int{
	"BAD_REQUEST":           400,
	"UNAUTHORIZED":          401,
	"FORBIDDEN":             403,
	"NOT_FOUND":             404,
	"INTERNAL_SERVER_ERROR": 500,
	"DUPLICATED":            409,
	"UNPROCESSABLE_ENTITY":  422,
	"TOO_MANY_REQUESTS":     429,
	"SERVICE_UNAVAILABLE":   503,
	"NOT_IMPLEMENTED":       501,
}

func httpError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func BadRequest(c *gin.Context, message string) {
	httpError(c, errorCodes["BAD_REQUEST"], message)
}

func Unauthorized(c *gin.Context, message string) {
	httpError(c, errorCodes["UNAUTHORIZED"], message)
}

func Forbidden(c *gin.Context, message string) {
	httpError(c, errorCodes["FORBIDDEN"], message)
}

func NotFound(c *gin.Context, message string) {
	httpError(c, errorCodes["NOT_FOUND"], message)
}

func InternalServerError(c *gin.Context, message string) {
	httpError(c, errorCodes["INTERNAL_SERVER_ERROR"], message)
}

func Duplicated(c *gin.Context, message string) {
	httpError(c, errorCodes["DUPLICATED"], message)
}

func UnprocessableEntity(c *gin.Context, message string) {
	httpError(c, errorCodes["UNPROCESSABLE_ENTITY"], message)
}

func TooManyRequests(c *gin.Context, message string) {
	httpError(c, errorCodes["TOO_MANY_REQUESTS"], message)
}

func ServiceUnavailable(c *gin.Context, message string) {
	httpError(c, errorCodes["SERVICE_UNAVAILABLE"], message)
}