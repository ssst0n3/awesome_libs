package cipher

import (
	"encoding/hex"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestCipher_SetKey(t *testing.T) {
	key := []byte("key")
	cipher := Cipher{}
	cipher.SetKey(key)
	assert.Equal(t, key, cipher.key)
}

func TestCipher_GetKey(t *testing.T) {
	cipherKeyHex, err := ioutil.ReadFile(PathCipherKeyTest)
	awesomeError.CheckErr(err)
	assert.Equal(t, nil, err)
	cipherKey, err := hex.DecodeString(strings.TrimSpace(string(cipherKeyHex)))
	awesomeError.CheckErr(err)
	assert.Equal(t, nil, err)
	cipher := Cipher{}
	err = cipher.GetKey(PathCipherKeyTest)
	awesomeError.CheckErr(err)
	assert.Equal(t, nil, err)
	assert.Equal(t, cipherKey, cipher.key)
}

func TestCipher_Success(t *testing.T) {
	cipher := Cipher{}
	cipher.GetKey(PathCipherKeyTest)
	plainText := "plain"
	cipherText, err := cipher.Encrypt(plainText)
	assert.Equal(t, nil, err)
	plain, err := cipher.Decrypt(cipherText)
	assert.Equal(t, nil, err)
	assert.Equal(t, plainText, plain)
}
