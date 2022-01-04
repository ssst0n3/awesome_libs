package awesome_error

import (
	"errors"
	"testing"
)

func TestCheckDebug(t *testing.T) {
	CheckDebug(errors.New("debug"))
}

func TestCheckErr(t *testing.T) {
	CheckErr(errors.New("error"))
}
