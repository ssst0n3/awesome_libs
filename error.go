package awesome_libs

func CheckErr(err error) {
	if err != nil {
		Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}
