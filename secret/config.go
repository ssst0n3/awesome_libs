package secret

import (
	"errors"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/ssst0n3/awesome_libs/secret/errs"
	"os"
)

var DirSecret string

func CheckDirSecretValid(dirSecret string) (valid bool) {
	if fileInfo, err := os.Stat(dirSecret); !os.IsNotExist(err) {
		valid = fileInfo.IsDir()
	} else {
		awesome_error.CheckErr(err)
		return
	}
	// TODO: warning insecure permission
	return
}

func GetDirSecretFromEnv() (dirSecret string, err error) {
	dirSecret = os.Getenv(consts.EnvDirSecret)
	if len(dirSecret) == 0 {
		err = errs.EnvDirSecretEmpty
	} else {
		if !CheckDirSecretValid(dirSecret) {
			err = errors.New(fmt.Sprintf(consts.ErrorDirSecretNotValid, dirSecret))
		}
	}
	return
}

/*
Please make sure the security of variable prefer by yourself
*/
func CreateDefaultSecretDir(prefer string) (dirSecret string) {
	{
		/* create dir assigned by environment */
		if len(prefer) > 0 {
			dirSecret = prefer
			if err := os.MkdirAll(dirSecret, 0755); err != nil {
				awesome_error.CheckWarning(err)
			} else {
				return
			}
		}
	}
	{
		/* create dir under root */
		dirSecret = consts.PreferDirSecretRoot
		if err := os.MkdirAll(dirSecret, 0755); err != nil {
			awesome_error.CheckWarning(err)
		} else {
			return
		}
	}
	{
		/* create dir under process dir */
		dirSecret = consts.PreferDirSecretProcess
		if err := os.MkdirAll(dirSecret, 0755); err != nil {
			awesome_error.CheckFatal(err)
		} else {
			return
		}
	}

	// TODO: is there need to get absolute path of dirSecret
	return
}

func GetDirSecret() (dirSecret string) {
	var err error
	if dirSecret, err = GetDirSecretFromEnv(); err != nil {
		dirSecret = CreateDefaultSecretDir(dirSecret)
	}
	return
}
