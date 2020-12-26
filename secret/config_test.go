package secret

import (
	"github.com/ssst0n3/awesome_libs/secret/consts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCheckDirSecretValid(t *testing.T) {
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
		assert.NoError(t, os.RemoveAll(dir))
		dirSecret := CreateDefaultSecretDir(dir)
		assert.Equal(t, dir, dirSecret)
	}
	t.Run("empty", func(t *testing.T) {
		dirSecret := CreateDefaultSecretDir("")
		assert.Equal(t, "", dirSecret)
	})
}

func TestGetDirSecret(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		os.Clearenv()
		dirSecret := GetDirSecret()
		assert.Equal(t, ".", dirSecret)
	})
	t.Run("not exists", func(t *testing.T) {
		os.Clearenv()
		dir := "secret"
		assert.NoError(t, os.Setenv(consts.EnvDirSecret, dir))
		dirSecret := GetDirSecret()
		assert.Equal(t, dir, dirSecret)
		assert.NoError(t, os.Remove(dir))
	})
}
