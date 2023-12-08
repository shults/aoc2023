package day08

import (
	"bufio"
	"fmt"
	"io"
)

func newProgram(in io.Reader) Program {
	reader := bufio.NewReader(in)

	var gen *DirectionGenerator
	instructionMap := NewInstructionGenerator()

	for {
		line, isPrefix, err := reader.ReadLine()

		panicOnFalse(!isPrefix, "prefix is not expected")

		if err == io.EOF {
			break
		}

		panicOnError(err)

		if len(line) == 0 {
			continue
		}

		if gen == nil {
			generator := NewDirectionGenerator(line)
			gen = &generator
			continue
		}

		instructionMap.Add(line)
	}

	return Program{
		directionGenerator:   *gen,
		instructionGenerator: instructionMap,
	}
}

type Program struct {
	directionGenerator   DirectionGenerator
	instructionGenerator InstructionGenerator
}

func (p *Program) Part1(verbose bool) int {
	iters := 0
	for {
		iters++
		res := p.instructionGenerator.Next(p.directionGenerator.Next())

		if verbose && iters%1000_000 == 0 {
			fmt.Printf("iter=%d res=%s\n", iters, res)
		}

		//if verbose {
		//	fmt.Printf("%s\n", res)
		//}

		//if iters == 100 {
		//	panic("ofof")
		//}

		if res == finalInstruction {
			break
		}
	}

	return iters
}
