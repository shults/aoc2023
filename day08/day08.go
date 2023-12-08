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
	err := flagSet.Parse(args)
	panicOnError(err)

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

	part1 := program.Part1(*verbose)

	fmt.Printf("part1=%d\n", part1)
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
