package loggerHelper

import (
	"io"
	"log"
	"os"
)

type SLogger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

var Logger *SLogger

func NewLogger(p string) *SLogger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &SLogger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func InitLogger(p string) {
	Logger = NewLogger(p)
}

func GetLogger() *SLogger {
	return Logger
}

// Create Non-Formatted Logs
func (l *SLogger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}
func (l *SLogger) Info(v ...interface{}) {
	l.info.Println(v...)
}
func (l *SLogger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}
func (l *SLogger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enabled Logs
func (l *SLogger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}
func (l *SLogger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}
func (l *SLogger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}
func (l *SLogger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}