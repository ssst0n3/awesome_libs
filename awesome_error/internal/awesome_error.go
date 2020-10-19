package internal

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type AwesomeError struct {
	Logger *logrus.Logger
}

func (a AwesomeError) CheckDebug(err error) {
	if err != nil {
		a.Logger.Debugf("%+v\n", errors.Errorf(err.Error()))
	}
}

func (a AwesomeError) CheckErr(err error) {
	if err != nil {
		a.Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}

func (a AwesomeError) CheckWarning(err error) {
	if err != nil {
		a.Logger.Warnf("%+v\n", errors.Errorf(err.Error()))
	}
}

func (a AwesomeError) CheckFatal(err error) {
	if err != nil {
		a.Logger.Fatalf("%+v\n", errors.Errorf(err.Error()))
	}
}
