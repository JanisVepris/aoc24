package util

// Replace a rune at a given index in a string
func StrReplaceAt(s string, i int, r rune) string {
	result := []rune(s)
	result[i] = r

	return string(result)
}
