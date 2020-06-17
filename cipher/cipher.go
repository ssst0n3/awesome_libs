package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	"io"
	"io/ioutil"
	"strings"
)

const (
	PanicKeyMustExists = "the cipher's key must be exists"
)

type Cipher struct {
	key []byte
}

func (c *Cipher) MustKeyExists() {
	if len(c.key) == 0 {
		panic(PanicKeyMustExists)
	}
}

func (c *Cipher) SetKey(key []byte) {
	c.key = key
}

func (c *Cipher) GetKey(keyPath string) error {
	cipherKeyHex, err := ioutil.ReadFile(keyPath)
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	cipherKey, err := hex.DecodeString(strings.TrimSpace(string(cipherKeyHex)))
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}

	c.key = cipherKey
	return nil
}

func (c *Cipher) Encrypt(pt string) (string, error) {
	c.MustKeyExists()
	plainText := []byte(pt)
	nonce := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, nonce)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}

	block, err := aes.NewCipher(c.key)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}

	cipherText := aesGcm.Seal(nil, nonce, plainText, nil)
	return hex.EncodeToString(append(nonce, cipherText...)), nil
}

func (c *Cipher) Decrypt(ct string) (string, error) {
	c.MustKeyExists()
	cipherText, err := hex.DecodeString(ct)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}
	nonce := cipherText[:12]
	block, err := aes.NewCipher(c.key)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}

	plainText, err := aesGcm.Open(nil, nonce, cipherText[12:], nil)
	if err != nil {
		awesomeError.CheckErr(err)
		return "", err
	}

	return string(plainText), nil
}

var CommonCipher = Cipher{}
