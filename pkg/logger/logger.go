package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

//go:generate mockgen -package=mock -source=logger.go -destination=$MOCK_FOLDER/logger.go Logger

// Logger describres methods associated with a Logger instance.
type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

// NewLogger instanciates a logger instance.
func NewLogger(appName string) Logger {
	var log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.JSONFormatter{},
		Level:     logrus.InfoLevel,
	}

	entry := log.WithFields(logrus.Fields{
		"appname": appName,
	})

	return entry
}
