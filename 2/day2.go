package day2

import (
	"aoc/util"
	"fmt"
	"strings"
)

func Run() {
	lines := util.ReadFile("2/day2.txt")

	reports := make([][]int, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		for _, part := range parts {
			reports[i] = append(reports[i], util.StrToInt(part))
		}
	}

	util.Timed(func() { part1(reports) })
	util.Timed(func() { part2(reports) })
}

func part1(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		if isSafe(report, false) {
			safeReports++
		}
	}

	fmt.Println("Result 1: ", safeReports)
}

func part2(reports [][]int) {
	safeReports := 0

	for _, report := range reports {
		if isSafe(report, true) {
			safeReports++
		}
	}
	fmt.Println("Result 2: ", safeReports)
}

func isSafe(report []int, errorTolerance bool) bool {
	i, j := 0, 1
	reportLen := len(report)

	isDecreasing := report[0] > report[1]

	for j < reportLen {
		val1 := report[i]
		val2 := report[j]
		hasError := false

		diff := util.Abs(val1 - val2)

		if isDecreasing && val1 < val2 {
			hasError = true
		}

		if !isDecreasing && val1 > val2 {
			hasError = true
		}

		if diff < 1 || diff > 3 {
			hasError = true
		}

		if errorTolerance && hasError {
			variation1 := util.RemoveFromSlice(report, i)
			if isSafe(variation1, false) {
				return true
			}

			variation2 := util.RemoveFromSlice(report, j)

			if isSafe(variation2, false) {
				return true
			}

			if i != 0 {
				variation3 := util.RemoveFromSlice(report, 0)

				if isSafe(variation3, false) {
					return true
				}
			}
		}

		if hasError {
			return false
		}

		i++
		j++
	}

	return true
}
