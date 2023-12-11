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
	emptyRowOrColumnSize := flagSet.Int("s", 1_000_000, "part 2")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	var lines []string

	if len(*inputFile) > 0 {
		lines, err = tools.ReadLinesFileLines(*inputFile)
	} else {
		lines, err = tools.ReadLines(in)
	}

	tools.AssertNoError(err)

	universe := NewUniverse(lines)

	fmt.Printf("part1=%d\n", universe.calculate(*emptyRowOrColumnSize))
}
