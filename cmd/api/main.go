package main

import (
	"context"
	"log"

	"github.com/Tonny-Francis/api-base-golang/internal/api"
	"github.com/Tonny-Francis/api-base-golang/internal/container"
	"github.com/Tonny-Francis/api-base-golang/pkg/core/server"
)

func main() {
	// Create context
	ctx := context.Background()

	// Create container
	ctx, deps, err := container.New(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Run server
	server.Run(
		ctx,
		deps,
		api.Router(ctx, deps),
	)
}
