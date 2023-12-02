package main

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"flag"
	"fmt"
	"os"
)

var programs = map[int]func(){
	1: day01.Main,
	2: day02.Main,
}

func main() {
	var day int
	flag.IntVar(&day, "d", 0, "day number")
	flag.Parse()

	prog, ok := programs[day]

	if !ok {
		fmt.Printf("Day with %d not found\n", day)
		os.Exit(1)
	}

	prog()
}
