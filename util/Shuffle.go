package util

import (
	"math/rand"
	"slices"
)

func Shuffle[T any](s []T) []T {
	r := slices.Clone(s)

	rand.Shuffle(len(r), func(i, j int) { r[i], r[j] = r[j], r[i] })

	return r
}
