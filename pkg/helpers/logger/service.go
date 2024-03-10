package logger

import "github.com/sirupsen/logrus"

type Logger interface {
	InitLogger(mode string) *logrus.Logger
}

type DefaultLogger struct{}

func NewLoggerService() Logger {
	return &DefaultLogger{}
}
