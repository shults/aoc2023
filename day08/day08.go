package day08

import (
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
	panicOnError(err)

	if !*p1 && !*p2 {
		*p1 = true
		*p2 = true
	}

	var program Program

	if len(*inputFile) > 0 {
		file, err := os.Open(*inputFile)

		if err != nil {
			flagSet.Usage()
			fmt.Printf("Error ocurred: %s\n", err)
			os.Exit(1)
		}

		program = newProgram(file)
		err = file.Close()
		panicOnError(err)
	} else {
		program = newProgram(in)
	}

	if *p1 {
		part1 := program.Part1()
		fmt.Printf("part1=%d\n", part1)
	}

	if *p2 {
		part2 := program.Part2(*verbose)
		fmt.Printf("part2=%d\n", part2)
	}
}

func panicOnFalse(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
