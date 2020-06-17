package common

import (
	error2 "github.com/ssst0n3/awesome_libs/error"
	"log"
	"os"
	"strings"
)

const ProjectName = "awesome_libs"

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
		log.Fatal("index must be bigger than -1!")
	}
	ProjectDir = dir[:index+len(ProjectName)] + "/"
}

func InitCipher() {
	if err := CommonCipher.GetKey(PathCipherKeyTest); err != nil {
		panic(err)
	}
}
