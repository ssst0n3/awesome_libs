package error_logger

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func Init(logger *logrus.Logger) {
	Logger = logger
}
