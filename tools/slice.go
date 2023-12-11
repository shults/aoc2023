package tools

func SliceCount[T any](in []T, fn func(T) bool) int {
	res := 0

	for _, item := range in {
		if fn(item) {
			res++
		}
	}

	return res
}
