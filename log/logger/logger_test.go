package logger

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInitLogger(t *testing.T) {
	var buffer bytes.Buffer
	logger := InitLogger("test", &buffer)
	logger.Info("apple")
	assert.Equal(t, true, strings.Contains(buffer.String(), "apple"))
}
