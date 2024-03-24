package main

import (
	"context"
	"log"

	"github.com/Tonny-Francis/api-scaffold-golang/internal/api"
	"github.com/Tonny-Francis/api-scaffold-golang/internal/container"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/server"
)

func main() {
	ctx := context.Background()

	ctx, helpers, services, domains, err := container.New(ctx, "default")

	if err != nil {
		log.Fatal(err)
	}

	server.NewServerService().Run(
		ctx,
		helpers,
		services,
		domains,
		api.Router(ctx, helpers, services, domains),
	)
}
