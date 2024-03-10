package server

import (
	"context"

	"github.com/Tonny-Francis/api-base-golang/internal/container"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run(ctx context.Context, helpers *container.Helpers, services *container.Services, domains *container.Domains, router *gin.Engine)
}

type DefaultServer struct{}

func NewServerService() Server {
	return &DefaultServer{}
}
