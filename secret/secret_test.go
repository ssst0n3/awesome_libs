package secret

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateKey(t *testing.T) {
	filename := "test"
	key, err := GenerateKey(filename)
	assert.NoError(t, err)
	assert.Equal(t, 32, len(key))
	assert.NoError(t, os.Remove(KeyFilePath(filename)))
}

func TestLoadKey(t *testing.T) {
	filename := "test"
	keyInit, init, err := LoadKey(filename)
	assert.Equal(t, 32, len(keyInit))
	assert.Equal(t, true, init)
	assert.NoError(t, err)

	key, init, err := LoadKey(filename)
	assert.Equal(t, keyInit, key)
	assert.Equal(t, false, init)
	assert.NoError(t, err)

	assert.NoError(t, os.Remove(KeyFilePath(filename)))
}
