package awesome_libs

import "github.com/pkg/errors"

func CheckErr(err error) {
	if err != nil {
		Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}

func CheckFatal(err error) {
	if err != nil {
		Logger.Fatalf("%+v\n", errors.Errorf(err.Error()))
	}
}
