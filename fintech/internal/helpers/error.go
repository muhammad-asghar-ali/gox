package helpers

// error handler
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
