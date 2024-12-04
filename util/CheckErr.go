package util

// Checks and error and panics if it is not nil
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
