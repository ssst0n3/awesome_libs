package cipher

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/awesome_libs/secret"
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"os"
)

const (
	FilenameCommonCipherKey = "awesome_libs_cipher_common"
	HintInitData            = "Hint: you need init your data, because the CommonCipher's key is init key."
)

func Init() {
	log.Logger.Info("cipher start init.")
	secret.InitDirSecret()
	var err error
	CommonCipher.key, IsInitKey, err = secret.LoadKey(FilenameCommonCipherKey)
	log.Logger.Info("cipher.IsInitKey:", IsInitKey)
	awesome_error.CheckFatal(err)
	if IsInitKey {
		log.Logger.Debug(HintInitData)
	}
}

func init() {
	if len(os.Getenv(consts.EnvDirSecret)) > 0 {
		Init()
	}
}
