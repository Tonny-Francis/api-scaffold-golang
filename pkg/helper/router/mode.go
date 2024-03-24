package router

import (
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/env"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func (r *DefaultRouter) InitGin(env *env.Env) *gin.Engine {
	gin.SetMode(env.GIN_MODE)

	router = gin.New()

	return router
}
