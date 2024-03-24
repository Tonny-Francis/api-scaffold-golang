package container

import (
	"context"

	"github.com/Tonny-Francis/api-scaffold-golang/pkg/domain/example"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/env"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/http"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/logger"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/router"
	"github.com/Tonny-Francis/api-scaffold-golang/pkg/helper/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

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
	logger := logger.NewLoggerService().InitLogger(env_mode)

	environments, err := env.NewEnvService().InitEnv(logger, env_mode)

	router := router.NewRouterService().InitGin(environments)

	if err != nil {
		return nil, nil, nil, nil, err
	}

	srvs := Services{}

	domains := Domains{
		Example: example.NewHandler(),
	}

	helpers := &Helpers{
		Logger:         logger,
		Env:            environments,
		Router:         router,
		HttpResponse:   http.NewResponseService(),
		HttpError:      http.NewErrorService(),
		RequestHandler: http.NewRequestHandler(),
		ParseSchema:    validator.NewSchemaValidatorService(),
	}

	logger.Info("Container initialized")

	return ctx, helpers, &srvs, &domains, nil
}
