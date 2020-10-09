package secret

import (
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCheckValid(t *testing.T) {
	assert.Equal(t, true, CheckDirSecretValid("/tmp"))
	assert.Equal(t, false, CheckDirSecretValid("/tmp/not_exists"))
}

func TestGetDirSecretFromEnv(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		dir := "/tmp"
		assert.NoError(t, os.Setenv(consts.EnvDirSecret, dir))
		dirSecret, err := GetDirSecretFromEnv()
		assert.NoError(t, err)
		assert.Equal(t, dir, dirSecret)
	})
}

func TestCreateDefaultSecretDir(t *testing.T) {
	{
		dir := "/tmp/secret"
		assert.NoError(t, os.Remove(dir))
		dirSecret := CreateDefaultSecretDir(dir)
		assert.Equal(t, dir, dirSecret)
	}
	t.Run("process cwd", func(t *testing.T) {
		assert.NoError(t, os.Remove(consts.PreferDirSecretProcess))
		dirSecret := CreateDefaultSecretDir("")
		assert.Equal(t, consts.PreferDirSecretProcess, dirSecret)
	})
}

func TestGetDirSecret(t *testing.T) {
	t.Run("process cwd", func(t *testing.T) {
		assert.NoError(t, os.Remove(consts.PreferDirSecretProcess))
		dirSecret := GetDirSecret()
		assert.Equal(t, consts.PreferDirSecretProcess, dirSecret)
	})
}
