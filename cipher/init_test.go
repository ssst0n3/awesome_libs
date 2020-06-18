package cipher

import (
	"errors"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"os"
	"strings"
)

const (
	ProjectName                 = "awesome_libs"
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
	awesome_error.CheckErr(err)
	index := strings.Index(dir, ProjectName)
	if index < 0 {
		awesome_error.CheckFatal(errors.New(PanicIndexMustBePositiveInt))
	}
	ProjectDir = dir[:index+len(ProjectName)] + "/"
}

func InitCipher() {
	if err := CommonCipher.GetKey(PathCipherKeyTest); err != nil {
		panic(err)
	}
}
