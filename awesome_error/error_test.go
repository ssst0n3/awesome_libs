package awesome_error

import (
	"errors"
	"testing"
)

func TestCheckErr(t *testing.T) {
	CheckErr(errors.New("test"))
}
