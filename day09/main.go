package day09

import (
	"aoc2023/tools"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// part1=1798691765
// part2=1104
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
	var sequences []Sequence

	for {
		line, isPrefix, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		tools.AssertTrue(!isPrefix, "prefix is not expected")

		seq, err := ParseSequence(string(line))
		tools.AssertNoError(err)

		sequences = append(sequences, *seq)
	}

	if *p1 {
		part1 := 0

		for _, seq := range sequences {
			part1 += seq.GetNext()
		}

		fmt.Printf("part1=%d\n", part1)
	}

	if *p2 {

		part2 := 0

		for _, seq := range sequences {
			part2 += seq.GetPrev()
		}

		fmt.Printf("part2=%d\n", part2)
	}
}
