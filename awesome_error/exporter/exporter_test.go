package exporter

import (
	"bytes"
	"errors"
	"github.com/ssst0n3/awesome_libs/log/logger"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetAwesomeError(t *testing.T) {
	var out bytes.Buffer
	a := GetAwesomeError(logger.InitLogger("my-logger", &out))
	a.CheckErr(errors.New("apple"))
	assert.Equal(t, true, strings.Contains(out.String(), "apple"))
}
