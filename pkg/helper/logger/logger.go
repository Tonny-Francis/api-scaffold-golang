package logger

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

type CustomTextFormatter struct {
	logrus.TextFormatter
}

var logger *logrus.Logger
var buffer *bytes.Buffer
var modeMutex sync.Mutex
var currentMode string

func (r *DefaultLogger) InitLogger(mode string) *logrus.Logger {
	modeMutex.Lock()
	
	defer modeMutex.Unlock()

	currentMode = mode

	logger = logrus.New()
	buffer = new(bytes.Buffer)

	logger.SetFormatter(&CustomTextFormatter{
		TextFormatter: logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "15:04:05.000",
			ForceColors:     true,
		},
	})

	logger.SetOutput(os.Stdout)

	return logger
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

	modeMutex.Lock()
	
	defer modeMutex.Unlock()

	if currentMode == "test" {
		buffer.WriteString(fmt.Sprintf("[%s] %s : %s\n", timestamp, level, message))
		return nil, nil
	}

	return []byte(fmt.Sprintf("[%s] %s : %s\n", timestamp, level, message)), nil
}