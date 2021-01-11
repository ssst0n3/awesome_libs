package awesome_structs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	A int `json:"a,string"`
}

func TestStringMap(t *testing.T) {
	result, err := StringMap(Test{A: 1})
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"a": "1"}, result)
}
