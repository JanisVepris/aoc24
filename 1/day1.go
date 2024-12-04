package day1

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadFile("1/day1.txt")

	list1 := []int{}
	list2 := []int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, "   ")
		list1 = append(list1, util.StrToInt(parts[0]))
		list2 = append(list2, util.StrToInt(parts[1]))
	}

	util.Timed(func() { part1(list1, list2) })
	util.Timed(func() { part2(list1, list2) })
}

func part1(list1, list2 []int) {
	distances := []int{}

	slices.Sort(list1)
	slices.Sort(list2)

	for i := range list1 {
		distances = append(distances, util.Abs(list1[i]-list2[i]))
	}

	sumFunc := func(value, carry int) int {
		return value + carry
	}

	result := util.Reduce(distances, sumFunc, 0)

	fmt.Println("Result 1: ", result)
}

func part2(list1, list2 []int) {
	similarityScore := 0

	occurences := map[int]int{}

	for _, value := range list1 {
		count := 0
		if _, ok := occurences[value]; !ok {
			for _, value2 := range list2 {
				if value == value2 {
					count++
				}
			}

			occurences[value] = count
		}

		similarityScore += value * occurences[value]
	}

	fmt.Println("Result 2: ", similarityScore)
}
