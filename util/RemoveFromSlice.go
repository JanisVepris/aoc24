package util

func RemoveFromSlice[T any](slice []T, index int) []T {
	copySlice := append([]T(nil), slice...)

	return append(copySlice[:index], copySlice[index+1:]...)
}
