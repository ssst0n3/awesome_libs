package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
)

func InitLogger(name string, writer io.Writer) *logrus.Logger {
	logger := logrus.New()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(writer)

	// Only log the warning severity or above.
	logger.SetLevel(logrus.InfoLevel)
	logger.Debug(fmt.Sprintf("%s's logger has been inited.", name))
	return logger
}
