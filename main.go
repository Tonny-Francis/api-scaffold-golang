package main

import (
	"github.com/Tonny-Francis/api-base-golang/configs/envConfig"
	"github.com/Tonny-Francis/api-base-golang/helpers/loggerHelper"
	"github.com/Tonny-Francis/api-base-golang/server"
)

func main() {
	// Initialize Logger
	loggerHelper.InitLogger("API")

	// Initialize Environment Configuration
	err := envConfig.InitEnvConfig()

	if err != nil {
		loggerHelper.Logger.Errorf("error initializing environment configuration: %v", err)
		return
	}
	
	// Initialize Routers
	server.InitRouters()
}
