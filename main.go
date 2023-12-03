package main

import (
	"aoc2023/day01"
	"aoc2023/day02"
	"aoc2023/day03"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var programs = map[int]func(set *flag.FlagSet, args []string, reader io.Reader){
	1: day01.Main,
	2: day02.Main,
	3: day03.Main,
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

	if len(errMsg) > 0 {
		fmt.Printf("Error: %s\n", errMsg)
	}

	os.Exit(1)
}
