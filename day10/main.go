package day10

import (
	"aoc2023/tools"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	verbose := flagSet.Bool("verbose", false, "verbose mode")
	inputFile := flagSet.String("f", "", "input file")
	p1 := flagSet.Bool("p1", false, "part 1")
	p2 := flagSet.Bool("p2", false, "part 2")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	if !*p1 && !*p2 {
		*p1 = true
		*p2 = true
	}

	if len(*inputFile) > 0 {
		file, err := os.Open(*inputFile)

		if err != nil {
			flagSet.Usage()
			fmt.Printf("Error ocurred: %s\n", err)
			os.Exit(1)
		}

		tools.AssertNoError(err)

		in = file
		defer file.Close()
	}

	reader := bufio.NewReader(in)
	var lines []string

	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		tools.AssertTrue(!isPrefix, "prefix is not expected")

		lines = append(lines, string(line))
	}

	graph := NewGraph(lines)

	if *p1 {
		part1 := graph.CalculatePart1()
		fmt.Printf("part1=%d\n", part1)
	}

	if *p2 {
		part2 := graph.CalculatePart2()

		fmt.Printf("part2=%d\n", part2)

		if *verbose {
			part22 := graph.CalculatePart2Analysis()
			fmt.Printf("part22=%d\n", part22)
		}
	}
}
