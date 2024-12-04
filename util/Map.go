package util

// Iterate over a slice and calls a function (fn) for each element, returns a slice of the results
func Map[T any, U any](slice []T, fn func(idx int, element T) U) (ret []U) {
	for i, item := range slice {
		ret = append(ret, fn(i, item))
	}

	return
}
