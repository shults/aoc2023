package day14

import (
	"aoc2023/tools"
	"flag"
	"fmt"
	"io"
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	verbose := flagSet.Bool("verbose", false, "verbose mode")
	inputFile := flagSet.String("f", "", "input file")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	var lines []string

	if len(*inputFile) > 0 {
		lines, err = tools.ReadLinesFileLines(*inputFile)
	} else {
		lines, err = tools.ReadLines(in)
	}

	tools.AssertNoError(err)
	plane := NewPlane(lines)

	fmt.Printf("part1=%d\n", plane.CalculatePart1(*verbose))
	fmt.Printf("part2=%d\n", plane.CalculatePart2(*verbose))
}
