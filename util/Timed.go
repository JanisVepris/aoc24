package util

import (
	"fmt"
	"time"
)

func Timed(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("Time: %s\n", elapsed)
}
