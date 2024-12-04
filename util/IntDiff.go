package util

import "math"

// IntDiff returns the absolute difference between two integers
func IntDiff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
