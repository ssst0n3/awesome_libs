package awesome_error

import (
	"github.com/pkg/errors"
	"github.com/ssst0n3/awesome_libs/log"
)

func CheckDebug(err error) {
	if err != nil {
		log.Logger.Debugf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckWarning(err error) {
	if err != nil {
		log.Logger.Warnf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckFatal(err error) {
	if err != nil {
		log.Logger.Fatalf("%+v\n", errors.Errorf(err.Error()))
	}
}
