package v1

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Routes(ctx context.Context, apiV1Router *gin.RouterGroup) {
	// Router health check
	apiV1Router.GET("/example", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}