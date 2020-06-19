package slice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIn(t *testing.T) {
	item := 1
	slice := []int{1, 2, 3}
	assert.Equal(t, true, In(item, slice))
}
