package util

import "strconv"

// Convert a numerical string to an integer
func StrToInt(numberString string) int {
	result, err := strconv.Atoi(numberString)

	CheckErr(err)

	return result
}
