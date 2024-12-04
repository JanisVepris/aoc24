package day3

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
)

func Run() {
	lines := util.ReadFile("3/day3.txt")
	input := ""

	for _, line := range lines {
		input += line
	}

	util.Timed(func() { part1(input) })
	util.Timed(func() { part2(input) })
}

func part1(input string) {
	instructions := getInstructions(input)

	result := util.Reduce(instructions, func(carry int, instruction []int) int {
		return carry + (instruction[0] * instruction[1])
	}, 0)

	fmt.Println("Result 1: ", result)
}

func part2(input string) {
	cleanString := parseDoInstructions(input)

	instructions := getInstructions(cleanString)

	result := util.Reduce(instructions, func(carry int, instruction []int) int {
		return carry + (instruction[0] * instruction[1])
	}, 0)

	fmt.Println("Result 2: ", result)
}

func getInstructions(input string) [][]int {
	regex := `mul\((\d{1,3}),(\d{1,3})\)`

	pattern := regexp.MustCompile(regex)

	result := pattern.FindAllStringSubmatch(input, -1)

	return util.Map(result, func(i int, instruction []string) []int {
		return parseInstruction(instruction)
	})
}

func parseInstruction(instruction []string) []int {
	return []int{util.StrToInt(instruction[1]), util.StrToInt(instruction[2])}
}

func parseDoInstructions(input string) string {
	remaining := strings.Clone(input)
	result := ""

	regex := regexp.MustCompile(`^mul\((\d{1,3}),(\d{1,3})\)`)
	instDo := "do()"
	instDont := "don't()"

	do := true

	for true {
		if len(remaining) == 0 {
			break
		}

		if strings.HasPrefix(remaining, instDo) {
			do = true
			remaining = strings.TrimPrefix(remaining, instDo)
			continue
		}

		if strings.HasPrefix(remaining, instDont) {
			do = false
			doIndex := strings.Index(remaining, instDo)

			if doIndex == -1 {
				break
			}

			remaining = remaining[doIndex:]

			continue
		}

		if !do {
			remaining = remaining[1:]
			continue
		}

		match := regex.FindString(remaining)

		if match != "" {
			result += match
			remaining = remaining[len(match):]
			continue
		}

		remaining = remaining[1:]
	}

	return result
}
