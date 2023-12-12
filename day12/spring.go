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

	arrangements = 0
	for _, line := range lines {
		arrangements += CalculateArrangementsPart2(line)
	}

	fmt.Printf("part2=%d\n", arrangements)
}

func CalculateArrangements(line string) int {
	parts := strings.Fields(line)
	tools.AssertTrue(len(parts) == 2, "expected 2 parts")
	symbols := parts[0]

	brokenLines, err := tools.Map(strings.Split(parts[1], ","), func(t *string) (int, error) {
		return strconv.Atoi(*t)
	})

	tools.AssertNoError(err)

	return calculateArrangementsOptimal(symbols, brokenLines)
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

func CalculateArrangementsPart2(str string) int {

	parts := strings.Fields(str)
	tools.AssertTrue(len(parts) == 2, "expected 2 parts")
	initPattern := parts[0]

	initMatchLine, err := tools.Map(strings.Split(parts[1], ","), func(t *string) (int, error) {
		return strconv.Atoi(*t)
	})

	tools.AssertNoError(err)

	const replacements = 5
	matchLine := make([]int, 0, len(initMatchLine)*replacements)

	patterns := []string{}
	for i := 0; i < replacements; i++ {
		matchLine = append(matchLine, initMatchLine...)
		patterns = append(patterns, initPattern)
	}

	pattern := strings.Join(patterns, "?")

	return calculateArrangementsOptimal(pattern, matchLine)
}

func extractMatchLine(matchLine string) (matches []int, hasNoMore bool, nextNum int) {
	matches = make([]int, 0)

	num := 0
	for i, symbol := range []byte(matchLine) {
		switch symbol {
		case unknown:
			//if num > 0 {
			//	matches = append(matches, num)
			//}

			// las is not full dont return result
			return matches, false, i
		case damaged:
			num++
		case operational:
			if num > 0 {
				matches = append(matches, num)
			}
			num = 0
		default:
			panic("unexpected symbol")
		}
	}

	if num > 0 {
		matches = append(matches, num)
	}

	return matches, true, 0
}

func calculateArrangementsOptimal(pattern string, matchLine []int) int {
	extracted, hasNoMore, nextUnknown := extractMatchLine(pattern)

	for i, minLength := 0, tools.Min(len(extracted), len(matchLine)); i < minLength; i++ {
		if matchLine[i] != extracted[i] {
			return 0
		}
	}

	if hasNoMore {
		if compareCombinations(extracted, matchLine) {
			return 1
		} else {
			return 0
		}
	} else {

		a := calculateArrangementsOptimal(
			pattern[:nextUnknown]+"#"+pattern[nextUnknown+1:],
			matchLine,
		)

		b := calculateArrangementsOptimal(
			pattern[:nextUnknown]+"."+pattern[nextUnknown+1:],
			matchLine,
		)

		return a + b
	}
}
