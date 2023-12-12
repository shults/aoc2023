package day12

import (
	"aoc2023/tools"
	"flag"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	operational byte = '.'
	damaged          = '#'
	unknown          = '?'
)

func Main(flagSet *flag.FlagSet, args []string, in io.Reader) {
	//verbose := flagSet.Bool("verbose", false, "verbose mode")
	inputFile := flagSet.String("f", "", "input file")
	err := flagSet.Parse(args)
	tools.AssertNoError(err)

	var lines []string

	if len(*inputFile) > 0 {
		lines, err = tools.ReadLinesFileLines(*inputFile)
	} else {
		lines, err = tools.ReadLines(in)
	}

	tools.AssertNoError(err)

	arrangements := 0

	for _, line := range lines {
		arrangements += CalculateArrangements(line)
	}

	fmt.Printf("part1=%d\n", arrangements)
}

func CalculateArrangements(line string) int {
	parts := strings.Fields(line)
	tools.AssertTrue(len(parts) == 2, "expected 2 parts")
	symbols := parts[0]

	brokenLines, err := tools.Map(strings.Split(parts[1], ","), func(t *string) (int, error) {
		return strconv.Atoi(*t)
	})

	tools.AssertNoError(err)

	arrangements := 0

	for _, combination := range generateCombinations([]string{symbols}) {
		//fmt.Printf("c: %+v d: %+v x: %+v\n", combination, describeCombination(combination), brokenLines)

		if compareCombinations(brokenLines, describeCombination(combination)) {
			arrangements++
		}
	}

	return arrangements
}

func compareCombinations(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func describeCombination(combination string) []int {
	res := make([]int, 0)

	for _, val := range strings.FieldsFunc(combination, func(r rune) bool {
		return r == '.'
	}) {
		res = append(res, len(val))
	}

	return res
}

func generateCombinations(lines []string) []string {
	res := make([]string, 0)

nextWord:
	for _, line := range lines {

		for i, symbol := range []byte(line) {
			if symbol == unknown {
				res = append(res, generateCombinations([]string{
					line[:i] + "." + line[i+1:],
					line[:i] + "#" + line[i+1:],
				})...)
				continue nextWord
			}
		}

		res = append(res, line)
	}
	return res
}

// the size of each contiguous group of damaged springs
