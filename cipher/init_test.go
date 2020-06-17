package cipher

import (
	"errors"
	error2 "github.com/ssst0n3/awesome_libs/error"
	"os"
	"strings"
)

const (
	ProjectName = "awesome_libs"
	PanicIndexMustBePositiveInt = "index must be bigger than -1"
)

var ProjectDir string
var PathCipherKeyTest string

func init() {
	GetProjectDir()
	PathCipherKeyTest = ProjectDir + "test_data/secret/cipher_key_test"
}

func GetProjectDir() {
	dir, err := os.Getwd()
	error2.CheckErr(err)
	index := strings.Index(dir, ProjectName)
	if index < 0 {
		error2.CheckFatal(errors.New(PanicIndexMustBePositiveInt))
	}
	ProjectDir = dir[:index+len(ProjectName)] + "/"
}

func InitCipher() {
	if err := CommonCipher.GetKey(PathCipherKeyTest); err != nil {
		panic(err)
	}
}
