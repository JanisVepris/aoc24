package util

import (
	"iter"
)

// Returns a block of strings from a slice of strings of defined size
// The block is returned as a slice of strings
// The iterator steps over the text left-to-right, top-to-bottom
// Each step advances the block by one column
// After the last column, the block is advanced by one row
// Hence it's "overlapping"
// This function assumes that the length of each row is equal
func StringOverlapBlockIter(s []string, rows, cols int) iter.Seq[[]string] {
	maxR := len(s)

	maxC := 0
	if maxR > 0 {
		maxC = len(s[0])
	}

	maxRowBound := maxR - rows + 1
	maxColBound := maxC - cols + 1

	return func(yield func([]string) bool) {
		if rows > maxR || cols > maxC || rows < 1 || cols < 1 {
			return
		}

		block := make([]string, rows)

		for r := 0; r < maxRowBound; r++ {
			for c := 0; c < maxColBound; c++ {
				for i := 0; i < rows; i++ {
					block[i] = s[r+i][c : c+cols]
				}

				if !yield(block) {
					return
				}
			}
		}
	}
}
