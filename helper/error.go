package helper

func PanicIfError(err error) {
	if err != nil {
		logger.Error(err)
		panic(err)
	}
}