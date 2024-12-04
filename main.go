package main

import (
	d1 "aoc/1"
	d2 "aoc/2"
	d3 "aoc/3"
	d4 "aoc/4"
	d5 "aoc/5"
	"fmt"
	"os"
	"strconv"
)

func main() {
	days := []func(){
		d1.Run,
		d2.Run,
		d3.Run,
		d4.Run,
		d5.Run,
	}

	day := getDay()

	if day > len(days) {
		fmt.Printf("Day %d not implemented yet\n", day)
		os.Exit(1)
	}

	fmt.Printf("Running day %d\n", day)
	days[day-1]()
}

func getDay() int {
	if len(os.Args) < 2 || len(os.Args) > 2 {
		usage()
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic(err)
	}

	if day < 1 || day > 25 {
		usage()
		os.Exit(1)
	}

	return day
}

func usage() {
	fmt.Println("Usage: go run . <dayNumber 1-25>")
}
