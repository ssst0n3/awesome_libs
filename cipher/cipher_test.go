package cipher

import (
	"encoding/hex"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
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
	assert.NoError(t, err)
	cipherKey, err := hex.DecodeString(strings.TrimSpace(string(cipherKeyHex)))
	assert.NoError(t, err)
	cipher := Cipher{}
	assert.NoError(t, cipher.GetKey(PathCipherKeyTest))
	assert.Equal(t, cipherKey, cipher.key)
}

func TestCipher_Success(t *testing.T) {
	cipher := Cipher{}
	cipher.GetKey(PathCipherKeyTest)
	plainText := "plain"
	cipherText, err := cipher.Encrypt(plainText)
	assert.NoError(t, err)
	plain, err := cipher.Decrypt(cipherText)
	assert.NoError(t, err)
	assert.Equal(t, plainText, plain)
}

func TestCommonCipher(t *testing.T) {
	assert.NoError(t, os.Setenv(consts.EnvDirSecret, "/tmp/secret"))
	Init()
	assert.Equal(t, true, len(CommonCipher.key) > 0)
}
