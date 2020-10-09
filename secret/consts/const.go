package consts

const (
	EnvDirSecret = "DIR_SECRET"
)

const (
	PreferDirSecretRoot    = "/secret"
	PreferDirSecretProcess = "./secret"
	PreferDirSecretTmp     = "/tmp/secret"
	PreferDirSecretCwd     = "."
)

const (
	ErrorEnvDirSecretEmpty = "environment variable ENV_SECRET is empty"
	ErrorDirSecretNotValid = "dir secret: %s is not valid"
)
