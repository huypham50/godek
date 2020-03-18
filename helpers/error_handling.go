package helpers

// Must -> error handing with Panic
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
