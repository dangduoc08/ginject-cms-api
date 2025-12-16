package slice

func ForEach[T any](arr []T, cb func(el T, i int)) {
	for i, el := range arr {
		cb(el, i)
	}
}
