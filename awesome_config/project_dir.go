package awesome_config

import (
	"os"
	"strings"
)

func GetProjectDir(ProjectName string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	index := strings.Index(dir, ProjectName)
	if index < 0 {
		panic("index must be bigger than -1!")
	}
	ProjectDir := dir[:index+len(ProjectName)]
	return ProjectDir
}
