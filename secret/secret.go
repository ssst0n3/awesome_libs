package secret

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"os"
	"path/filepath"
)

func KeyFilePath(filename string) string {
	return DirSecret + string(filepath.Separator) + filename
}

func CheckSecretFileValid(filename string) (valid bool) {
	path := KeyFilePath(filename)
	if fileInfo, err := os.Stat(path); !os.IsNotExist(err) {
		if !fileInfo.IsDir() {
			valid = true
		}
	}
	return
}

func GenerateKey(filename string) (key []byte, err error) {
	key = make([]byte, 32)

	path := KeyFilePath(filename)
	b := make([]byte, 16)
	if _, err = rand.Read(b); err != nil {
		awesome_error.CheckErr(err)
		return
	}
	_ = hex.Encode(key, b)
	if err = ioutil.WriteFile(path, key, 0600); err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

/*
* filename: under secret dir
* you can assign secret dir by environment variable DIR_SECRET
 */
func LoadKey(filename string) (key []byte, init bool, err error) {
	if !CheckSecretFileValid(filename) {
		/* init */
		init = true
		key, err = GenerateKey(filename)
	} else {
		/* load old key */
		key, err = ioutil.ReadFile(KeyFilePath(filename))
	}
	return
}
