package slice

func Map[T, U any](arr []T, cb func(el T, i int) U) []U {
	newArr := make([]U, len(arr))
	ForEach(arr, func(el T, i int) {
		newArr[i] = cb(el, i)
	})

	return newArr
}
