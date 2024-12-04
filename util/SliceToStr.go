package util

import (
	"fmt"
	"strings"
)

// SpliceToStr converts a slice to a string with a delimiter
func SpliceToStr[T any](slice []T, delimiter string) string {
	return strings.Trim(
		strings.Join(
			strings.Fields(
				fmt.Sprintf("%v", slice),
			),
			delimiter,
		),
		"[]",
	)
}
