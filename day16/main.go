package day16

import (
	"aoc2023/tools"
	"flag"
	"fmt"
	"io"
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	inputFile := flagSet.String("f", "", "input file")
	verbose := flagSet.Bool("verbose", false, "input file")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	var lines []string
	if len(*inputFile) > 0 {
		lines, err = tools.ReadLinesFileLines(*inputFile)
	} else {
		lines, err = tools.ReadLines(in)
	}

	tools.AssertNoError(err)
	beamSplitter := NewBeamSplitter(lines)

	fmt.Printf("part1=%d\n", beamSplitter.CalculateEnergizedTiles(*verbose))
	//fmt.Printf("part2=%d\n", beamSplitter.CalculateFocalSum())
}
