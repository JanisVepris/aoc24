package util

import (
	"bufio"
	"os"
)

// Reads a file and returns a slice of its lines as strings
func ReadFile(filename string) (lines []string) {
	wd, err := os.Getwd()
	CheckErr(err)

	file, err := os.Open(wd + "/" + filename)
	CheckErr(err)

	buffer := bufio.NewScanner(file)

	for buffer.Scan() {
		lines = append(lines, buffer.Text())
	}

	file.Close()

	return lines
}
