package tools

func Every[T any](items []T, predicate func(T) bool) bool {
	for _, item := range items {
		if predicate(item) == false {
			return false
		}
	}

	return true
}
