package logrus_helper

import (
	"github.com/sirupsen/logrus"
	"io"
)

func NewLogrus(writer io.Writer) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	if writer != nil {
		logger.SetOutput(writer)
	}
	logger.SetLevel(logrus.InfoLevel)
	logger.Info()
	return logger
}
