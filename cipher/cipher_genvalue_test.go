/*
just for encrypting and decrypting manually
*/
package cipher

import (
	"encoding/hex"
	"github.com/ssst0n3/awesome_libs/log"
	"testing"
)

func TestCipherEncrypt(t *testing.T) {
	key, _ := hex.DecodeString("92392daec439e38014c8d697014034a4b175ad8fcb16471328f90f3c1493c573")
	cipher := Cipher{
		[]byte(key),
	}
	plainText := []byte("qwxf{you_say_chick_beautiful?}")
	plainText = []byte("")
	plainText = []byte("user")
	cipherText, _ := cipher.Encrypt(plainText)
	log.Logger.Info(cipherText)
}

func TestCipherDecrypt(t *testing.T) {
	key, _ := hex.DecodeString("92392daec439e38014c8d697014034a4b175ad8fcb16471328f90f3c1493c573")
	cipher := Cipher{
		[]byte(key),
	}
	cipherText := "17903a209efbfee2716d252e847507c90f060beb19b426d6b250c2de"
	plainText, _ := cipher.Decrypt(cipherText)
	log.Logger.Info(plainText)
}
