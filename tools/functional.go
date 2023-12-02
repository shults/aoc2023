package tools

func MustEvery[T any](items []T, predicate func(*T) bool) bool {
	for _, item := range items {
		if predicate(&item) == false {
			return false
		}
	}

	return true
}

func MustEach[T any](items []T, mapFn func(*T)) {
	for _, item := range items {
		mapFn(&item)
	}
}

func Each[T any](items []T, mapFn func(*T) error) error {
	for _, item := range items {
		if err := mapFn(&item); err != nil {
			return err
		}
	}
	return nil
}

func MustMap[T any, R any](items []T, mapFn func(*T) R) []R {
	res := make([]R, 0, cap(items))

	for _, item := range items {
		res = append(res, mapFn(&item))
	}

	return res
}

func Map[T any, R any](items []T, mapFn func(*T) (R, error)) ([]R, error) {
	res := make([]R, 0, cap(items))

	for _, item := range items {
		mapped, err := mapFn(&item)

		if err != nil {
			return nil, err
		}

		res = append(res, mapped)
	}

	return res, nil
}
