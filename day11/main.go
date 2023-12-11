package day11

import (
	"aoc2023/tools"
	"flag"
	"fmt"
	"io"
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	//verbose := flagSet.Bool("verbose", false, "verbose mode")
	inputFile := flagSet.String("f", "", "input file")
	p1 := flagSet.Bool("p1", false, "part 1")
	p2 := flagSet.Bool("p2", false, "part 2")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	if !*p1 && !*p2 {
		*p1 = true
		*p2 = true
	}

	var lines []string

	if len(*inputFile) > 0 {
		lines, err = tools.ReadLinesFileLines(*inputFile)
	} else {
		lines, err = tools.ReadLines(in)
	}

	tools.AssertNoError(err)

	universe := NewUniverse(lines)

	if *p1 {
		part1 := universe.Part1()
		fmt.Printf("part1=%d\n", part1)
	}

	if *p2 {
		part2 := universe.Part1()
		fmt.Printf("part2=%d\n", part2)
	}
}
