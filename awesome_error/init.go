package awesome_error

import (
	"github.com/ssst0n3/awesome_libs/awesome_error/error_logger"
	"github.com/ssst0n3/awesome_libs/log"
)

func init() {
	error_logger.Init(log.Logger)
}
