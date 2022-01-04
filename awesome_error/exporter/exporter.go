package exporter

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error/internal"
)

func GetAwesomeError(logger *logrus.Logger, details bool) internal.AwesomeError {
	return internal.AwesomeError{
		Details: details,
		Logger:  logger,
	}
}
