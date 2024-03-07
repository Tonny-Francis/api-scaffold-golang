package container

import (
	"context"

	"github.com/Tonny-Francis/api-base-golang/pkg/core/env"
	"github.com/Tonny-Francis/api-base-golang/pkg/core/logger"
	"github.com/Tonny-Francis/api-base-golang/pkg/core/router"
	"github.com/Tonny-Francis/api-base-golang/pkg/domain/example"
	"github.com/Tonny-Francis/api-base-golang/pkg/services/http"
	"github.com/Tonny-Francis/api-base-golang/pkg/services/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type components struct {
	Logger *logrus.Logger
	Env    *env.Env
	Router *gin.Engine
}

type Services struct {
	HttpResponse   http.Response
	HttpError      http.Error
	RequestHandler http.Handler
	ParseSchema    validator.SchemaValidator
	Example        example.Handler
}

type Dependencies struct {
	Components components
	Services   Services
}

func New(ctx context.Context, env_mode string) (context.Context, *Dependencies, error) {
	cmps, err := setupComponents(ctx, env_mode)

	if err != nil {
		return nil, nil, err
	}

	srvs := Services{
		HttpResponse:   http.NewResponseService(),
		HttpError:      http.NewErrorService(),
		RequestHandler: http.NewRequestHandler(),
		ParseSchema:    validator.NewSchemaValidatorService(),
		Example:        example.NewHandler(),
	}

	deps := Dependencies{
		Components: *cmps,
		Services:   srvs,
	}

	return ctx, &deps, nil
}

func setupComponents(ctx context.Context, env_mode string) (*components, error) {
	logger := logger.InitLogger(env_mode)

	env, err := env.InitEnv(logger, env_mode)

	if err != nil {
		return nil, err
	}

	router := router.InitGin(env)

	return &components{
		Logger: logger,
		Env:    env,
		Router: router,
	}, nil
}
