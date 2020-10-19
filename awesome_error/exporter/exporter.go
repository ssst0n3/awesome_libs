package exporter

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error/internal"
)

func GetAwesomeError(logger *logrus.Logger) internal.AwesomeError {
	return internal.AwesomeError{Logger: logger}
}
