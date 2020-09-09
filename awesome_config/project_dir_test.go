package awesome_config

import (
	"github.com/ssst0n3/awesome_libs/log"
	"testing"
)

func TestGetProjectDir(t *testing.T) {
	log.Logger.Debug(GetProjectDir("awesome_libs"))
}
