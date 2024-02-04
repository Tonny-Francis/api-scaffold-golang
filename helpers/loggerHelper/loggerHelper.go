package loggerHelper

import (
	"fmt"
	"os"
	"strings"
	"github.com/sirupsen/logrus"
	"github.com/logrusorgru/aurora"
)

type CustomTextFormatter struct {
	logrus.TextFormatter
}

func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format(f.TimestampFormat)
	level := strings.ToUpper(entry.Level.String())
	message := entry.Message

	switch entry.Level {
	case logrus.InfoLevel:
		level = aurora.BrightGreen(level).String()
	case logrus.WarnLevel:
		level = aurora.BrightYellow(level).String()
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		level = aurora.BrightRed(level).String()
	}

	return []byte(fmt.Sprintf("[%s] %s : %s\n", timestamp, level, message)), nil
}

var Logger *logrus.Logger

func InitLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&CustomTextFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "15:04:05.000",
			ForceColors:     true,
		},
	})

	Logger.SetOutput(os.Stdout)
}

func GetLogger() *logrus.Logger {
	return Logger
}