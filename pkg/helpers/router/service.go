package router

import (
	"github.com/Tonny-Francis/api-base-golang/pkg/helpers/env"
	"github.com/gin-gonic/gin"
)

type Router interface {
	InitGin(env *env.Env) *gin.Engine
}

type DefaultRouter struct{}

func NewRouterService() Router {
	return &DefaultRouter{}
}
