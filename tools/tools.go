package tools

func AssertTrue(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func AssertNoError(err error) {
	if err != nil {
		panic(err)
	}
}
