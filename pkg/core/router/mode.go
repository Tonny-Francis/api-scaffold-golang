package router

import (
	"github.com/Tonny-Francis/api-base-golang/pkg/core/env"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitGin(env *env.Env) *gin.Engine {
	gin.SetMode(env.GIN_MODE)

	router = gin.New()

	return router
}