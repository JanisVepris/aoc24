package day4

import (
	"aoc/util"
	"fmt"
	"slices"
)

func Run() {
	lines := util.ReadFile("4/day4.txt")

	util.Timed(func() { part1(lines) })
	util.Timed(func() { part2(lines) })
}

func part1(input []string) {
	result := 0

	for y, line := range input {
		for x := 0; x < len(line); x++ {
			result += findXmas1(input, x, y)
		}
	}

	fmt.Println("Result 1: ", result)
}

func part2(input []string) {
	result := 0

	for block := range util.StringOverlapBlockIter(input, 3, 3) {
		result += finxXmas2(block)
	}

	fmt.Println("Result 2: ", result)
}

func finxXmas2(block []string) int {
	if string(block[1][1]) != "A" {
		return 0
	}

	word1 := string(block[0][0]) + string(block[2][2])
	word2 := string(block[0][2]) + string(block[2][0])

	words := []string{"MS", "SM"}

	if !slices.Contains(words, word1) {
		return 0
	}

	if !slices.Contains(words, word2) {
		return 0
	}

	return 1
}

func findXmas1(lines []string, posX, posY int) int {
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	found := 0
	word := ""
	needle := "XMAS"

	if lines[posY][posX] != 'X' {
		return 0
	}

	//search left
	if posX >= 3 {
		word = util.StrReverse(lines[posY][posX-3 : posX+1])

		if word == needle {
			found++
		}
	}

	// search right
	if posX <= maxX-3 {
		word = lines[posY][posX : posX+4]

		if word == needle {
			found++
		}
	}

	// search top
	if posY >= 3 {
		word = ""

		for i := 0; i < 4; i++ {
			word += string(lines[posY-i][posX])
		}

		if word == needle {
			found++
		}
	}

	// search bottom
	if posY <= maxY-3 {
		word = ""

		for i := 0; i < 4; i++ {
			word += string(lines[posY+i][posX])
		}

		if word == needle {
			found++
		}
	}

	// search diagonal top left
	if posX >= 3 && posY >= 3 {
		word = ""

		for i := 0; i < 4; i++ {
			word += string(lines[posY-i][posX-i])
		}

		if word == needle {
			found++
		}
	}

	// search diagonal top right
	if posX <= maxX-3 && posY >= 3 {
		word = ""

		for i := 0; i < 4; i++ {
			word += string(lines[posY-i][posX+i])
		}

		if word == needle {
			found++
		}
	}

	// search diagonal bottom left
	if posX >= 3 && posY <= maxY-3 {
		word = ""

		for i := 0; i < 4; i++ {
			word += string(lines[posY+i][posX-i])
		}

		if word == needle {
			found++
		}
	}

	// search diagonal bottom right
	if posX <= maxX-3 && posY <= maxY-3 {
		word = ""

		for i := 0; i < 4; i++ {
			word += string(lines[posY+i][posX+i])
		}

		if word == needle {
			found++
		}
	}

	return found
}
