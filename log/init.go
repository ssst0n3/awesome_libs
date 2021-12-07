package log

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log/logger"
	"os"
)

var Logger *logrus.Logger
var File *os.File

func init() {
	Logger = logger.InitLogger("awesome_libs", os.Stdout)
}

func Output2File(filename string) (err error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	File = file
	Logger.Out = file
	return
}

func CloseFile() error {
	return File.Close()
}
