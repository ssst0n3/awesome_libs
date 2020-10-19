package awesome_error

import (
	"github.com/pkg/errors"
	"github.com/ssst0n3/awesome_libs/awesome_error/error_logger"
)

func CheckDebug(err error) {
	if err != nil {
		error_logger.Logger.Debugf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckErr(err error) {
	if err != nil {
		error_logger.Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckWarning(err error) {
	if err != nil {
		error_logger.Logger.Warnf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckFatal(err error) {
	if err != nil {
		error_logger.Logger.Fatalf("%+v\n", errors.Errorf(err.Error()))
	}
}
