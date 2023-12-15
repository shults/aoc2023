package day15

import (
	"aoc2023/tools"
	"flag"
	"fmt"
	"io"
	"os"
)

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

	fmt.Printf("part1=%d\n", HashSumPart1(in))
}
