package env

import "github.com/sirupsen/logrus"

type Environment interface {
	InitEnv(logger *logrus.Logger, mode string) (*Env, error)
}

type DefaultEnvironment struct{}

func NewEnvService() Environment {
	return &DefaultEnvironment{}
}
