package errs

import (
	"errors"
	"github.com/ssst0n3/awesome_libs/secret/consts"
)

var (
	EnvDirSecretEmpty = errors.New(consts.ErrorEnvDirSecretEmpty)
)
