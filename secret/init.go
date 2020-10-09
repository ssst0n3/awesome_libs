package secret

func InitDirSecret() {
	DirSecret = GetDirSecret()
}

func init() {
	InitDirSecret()
}
