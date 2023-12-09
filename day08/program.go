package day08

import (
	"aoc2023/tools"
	"bufio"
	"io"
)

func newProgram(in io.Reader) Program {
	reader := bufio.NewReader(in)

	var gen *DirectionGenerator

	var lines []string
	ctr := 0

	for {
		ctr++
		line, isPrefix, err := reader.ReadLine()

		tools.AssertTrue(!isPrefix, "prefix is not expected")

		if err == io.EOF {
			break
		}

		tools.AssertNoError(err)

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

func (p *Program) Part1() int {
	return p.instructionGenerator.Part1(p.directionGenerator)
}

func (p *Program) Part2(verbose bool) int {
	return p.instructionGenerator.Part2(p.directionGenerator, verbose)
}
