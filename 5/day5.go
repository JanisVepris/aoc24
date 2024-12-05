package day5

import (
	"aoc/util"
	"fmt"
	"slices"
	"strings"
)

func Run() {
	lines := util.ReadFile("5/day5.txt")

	rules := map[string][]string{}
	updates := [][]string{}

	isRules := true

	for _, line := range lines {
		if line == "" {
			isRules = false
			continue
		}

		if isRules {
			parts := strings.Split(line, "|")
			if rule, ok := rules[parts[0]]; ok {
				rules[parts[0]] = append(rule, parts[1])
			} else {
				rules[parts[0]] = []string{parts[1]}
			}
		} else {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	invalidUpdates := [][]string{}
	util.Timed(func() { invalidUpdates = part1(rules, updates) })
	util.Timed(func() { part2(rules, invalidUpdates) })
}

func part1(rules map[string][]string, updates [][]string) [][]string {
	result := 0

	invalidUpdates := [][]string{}

	for _, update := range updates {
		if isValid(rules, update) {
			updateLen := len(update)
			result += util.StrToInt(update[updateLen/2])
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	fmt.Println("Result 1: ", result)

	return invalidUpdates
}

func part2(rules map[string][]string, updates [][]string) {
	result := 0

	for _, update := range updates {
		for !isValid(rules, update) {
			update = swapInvalidPages(rules, update)
		}

		result += util.StrToInt(update[len(update)/2])
	}

	fmt.Println("Result 2: ", result)
}

func swapInvalidPages(rules map[string][]string, update []string) []string {
	for idx, page := range update {
		pageRules, ok := rules[page]

		if !ok {
			continue
		}

		for _, rule := range pageRules {
			ruleIdx := slices.Index(update, rule)

			if ruleIdx == -1 {
				continue
			}

			if ruleIdx < idx {
				update = swapUpdates(update, idx, ruleIdx)
			}
		}
	}

	return update
}

func isValid(rules map[string][]string, update []string) bool {
	for idx, page := range update {
		pageRules, ok := rules[page]

		if !ok {
			continue
		}

		for _, rule := range pageRules {
			ruleIdx := slices.Index(update, rule)

			if ruleIdx == -1 {
				continue
			}

			if ruleIdx < idx {
				return false
			}
		}
	}

	return true
}

func swapUpdates(update []string, idx1, idx2 int) []string {
	update[idx1], update[idx2] = update[idx2], update[idx1]
	return update
}
