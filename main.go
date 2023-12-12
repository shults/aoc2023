package main

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"aoc2023/day03"
	"aoc2023/day04"
	"aoc2023/day05"
	"aoc2023/day06"
	"aoc2023/day07"
	"aoc2023/day08"
	"aoc2023/day09"
	"aoc2023/day10"
	"aoc2023/day11"
	"aoc2023/day12"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var programs = map[int]func(set *flag.FlagSet, args []string, reader io.Reader){
	1:  day01.Main,
	2:  day02.Main,
	3:  day03.Main,
	4:  day04.Main,
	5:  day05.Main,
	6:  day06.Main,
	7:  day07.Main,
	8:  day08.Main,
	9:  day09.Main,
	10: day10.Main,
	11: day11.Main,
	12: day12.Main,
}

func main() {
	if len(os.Args) == 1 {
		printUsageAndExit("")
	}

	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		printUsageAndExit("")
	}

	prog, ok := programs[day]

	if !ok {
		printUsageAndExit(
			fmt.Sprintf("day with %d not found\n", day),
		)
	}

	flagSet := flag.NewFlagSet(fmt.Sprintf("%s %d [opts]", os.Args[0], day), flag.ExitOnError)
	prog(flagSet, os.Args[2:], os.Stdin)
}

func printUsageAndExit(errMsg string) {
	fmt.Printf("Usage:\n\t%s <day> [opts]\n", os.Args[0])
	fmt.Printf("\t%s help <day>\n", os.Args[0])

	if len(errMsg) > 0 {
		fmt.Printf("Error: %s\n", errMsg)
	}

	os.Exit(1)
}
