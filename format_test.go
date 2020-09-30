package awesome_libs

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestFormat(t *testing.T) {
	msg := Format(
		"Hello {{.name}}", Dict{
			"name": "awesome",
		},
	)
	assert.Equal(t, "Hello awesome", msg)
}
