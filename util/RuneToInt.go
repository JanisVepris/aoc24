package util

import "strconv"

// Convert a rune to an integer
func RuneToInt(value rune) int {
	result, err := strconv.Atoi(string(value))

	CheckErr(err)

	return result
}
