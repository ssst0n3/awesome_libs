package exporter

import (
	"bytes"
	"errors"
	"github.com/ssst0n3/awesome_libs/log/logger"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestGetAwesomeError(t *testing.T) {
	var out bytes.Buffer
	a := GetAwesomeError(logger.InitLogger("my-logger", &out), true)
	a.CheckErr(errors.New("apple"))
	assert.Equal(t, true, strings.Contains(out.String(), "apple"))
}

func ExampleGetAwesomeError() {
	ae := GetAwesomeError(logger.InitLogger("my-logger", os.Stdout), true)
	ae.CheckErr(errors.New("apple"))
}
