package day15

import (
	"aoc2023/tools"
	"flag"
	"fmt"
	"io"
	"os"
)

// part1=505427
// part2=243747
func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	inputFile := flagSet.String("f", "", "input file")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	if len(*inputFile) > 0 {
		file, err := os.Open(*inputFile)
		tools.AssertNoError(err)
		defer file.Close()
		in = file
	}

	boxes := NewBoxes()
	boxes.ProcessInput(in)

	fmt.Printf("part1=%d\n", boxes.CalculateSumOfHashes())
	fmt.Printf("part2=%d\n", boxes.CalculateFocalSum())
}
