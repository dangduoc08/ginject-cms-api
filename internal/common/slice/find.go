package slice

func Find[T any](arr []T, cb func(el T, i int) bool) T {
	for i, el := range arr {
		if cb(el, i) {
			return el
		}
	}

	var zero T
	return zero
}
