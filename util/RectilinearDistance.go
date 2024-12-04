package util

import "aoc/util/types"

// Calculates the rectilinear distance between two points
// The distance is define as the sum of the absolute differences of Cartesian coordinates
//
// Also known as Manhattan distance or Taxicab distance
func RectilinearDistance(a, b types.Point2D) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}
