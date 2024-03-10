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
}

type Domains struct {
	Example example.Handler
}

type Helpers struct {
	Logger         *logrus.Logger
	Env            *env.Env
	Router         *gin.Engine
	HttpResponse   http.Response
	HttpError      http.Error
	RequestHandler http.Handler
	ParseSchema    validator.SchemaValidator
}

func New(ctx context.Context, env_mode string) (context.Context, *Helpers, *Services, *Domains, error) {
	components, err := setupComponents(ctx, env_mode)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	srvs := Services{}

	domains := Domains{
		Example: example.NewHandler(),
	}

	helpers := &Helpers{
		Logger:         components.Logger,
		Env:            components.Env,
		Router:         components.Router,
		HttpResponse:   http.NewResponseService(),
		HttpError:      http.NewErrorService(),
		RequestHandler: http.NewRequestHandler(),
		ParseSchema:    validator.NewSchemaValidatorService(),
	}

	return ctx, helpers, &srvs, &domains, nil
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
