package tools

import "cmp"

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

func Min[T cmp.Ordered](a, b T) T {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
