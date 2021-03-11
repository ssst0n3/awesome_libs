package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutput2File(t *testing.T) {
	assert.NoError(t, Output2File("/tmp/tmp"))
	Logger.Info("test")
}