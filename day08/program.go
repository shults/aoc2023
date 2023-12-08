package day08

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"time"
)

func newProgram(in io.Reader) Program {
	reader := bufio.NewReader(in)

	var gen *DirectionGenerator

	var lines []string
	ctr := 0

	for {
		ctr++
		line, isPrefix, err := reader.ReadLine()

		panicOnFalse(!isPrefix, "prefix is not expected")

		if err == io.EOF {
			break
		}

		panicOnError(err)

		if ctr == 1 {
			generator := NewDirectionGenerator(line)
			gen = &generator
			continue
		}

		if ctr == 2 {
			continue
		}

		lines = append(lines, string(line))
	}

	instructionMap := NewInstructionGenerator(lines)

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

	offset := 30
	bitValue := int(0)
	mask := 1 << offset

	for {
		iters++

		if verbose {
			newBitValue := (iters & mask) >> offset

			if newBitValue != bitValue {
				fmt.Printf("iters=%d percent=%f time=%s\n", iters, float64(iters)/float64(math.MaxInt64), time.Now())
			}

			bitValue = newBitValue
		}

		res := p.instructionGenerator.Next(p.directionGenerator.Next())

		if res == p.instructionGenerator.finalInstruction {
			break
		}
	}

	return iters
}

func (p *Program) Part2(verbose bool) int {
	return p.instructionGenerator.Part2(p.directionGenerator, verbose)
}
