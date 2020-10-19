package awesome_error

import (
	"github.com/ssst0n3/awesome_libs/awesome_error/exporter"
	"github.com/ssst0n3/awesome_libs/log"
)

var Default = exporter.GetAwesomeError(log.Logger)

func CheckDebug(err error) {
	Default.CheckDebug(err)
}

func CheckErr(err error) {
	Default.CheckErr(err)
}

func CheckWarning(err error) {
	Default.CheckWarning(err)
}

func CheckFatal(err error) {
	Default.CheckFatal(err)
}
