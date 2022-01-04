package internal

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AwesomeError struct {
	Details bool
	Logger  *logrus.Logger
}

func (a AwesomeError) CheckDebug(err error) {
	if err != nil {
		if a.Details {
			a.Logger.Debugf("%+v\n", errors.Errorf(err.Error()))
		} else {
			a.Logger.Debug(err)
		}
	}
}

func (a AwesomeError) CheckErr(err error) {
	if err != nil {
		if a.Details {
			a.Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
		} else {
			a.Logger.Error(err)
		}
	}
}

func (a AwesomeError) CheckWarning(err error) {
	if err != nil {
		if a.Details {
			a.Logger.Warnf("%+v\n", errors.Errorf(err.Error()))
		} else {
			a.Logger.Warn(err)
		}
	}
}

func (a AwesomeError) CheckFatal(err error) {
	if err != nil {
		if a.Details {
			a.Logger.Fatalf("%+v\n", errors.Errorf(err.Error()))
		} else {
			a.Logger.Fatal(err)
		}
	}
}
