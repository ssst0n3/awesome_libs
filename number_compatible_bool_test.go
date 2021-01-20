package awesome_libs

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberCompatibleBool_UnmarshalJSON(t *testing.T) {
	var b NumberCompatibleBool
	err := json.Unmarshal([]byte(`1`), &b)
	assert.NoError(t, err)
	assert.Equal(t, true, bool(b))
}
