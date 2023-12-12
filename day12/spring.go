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

// 74262028574941 -> too low

func CalculateArrangementsPart2(str string) int {
	return CalculateArrangementsPart2parametrized(str, 5)
}

// / CalculateArrangementsPart2parametrized Deprecated
func CalculateArrangementsPart2parametrized(str string, times int) int {
	// it doesn't work

	r1 := calculateArrangementsPart2(str, 1)
	r2 := calculateArrangementsPart2(str, 2)
	//r3 := calculateArrangementsPart2(str, 3)

	multiplier := r2 / r1
	rest := r2 % r1
	res := r1

	fmt.Printf("[%s] => r1=%d r2=%d mp=%d res=%d rest=%d\n", str, r1, r2, multiplier, res, rest)

	if rest != 0 {
		// full calc
		return calculateArrangementsPart2(str, times)
	}

	for i := 1; i < times; i++ {
		res *= multiplier
	}

	return res
}

func calculateArrangementsPart2(str string, times int) int {

	parts := strings.Fields(str)
	tools.AssertTrue(len(parts) == 2, "expected 2 parts")
	initPattern := parts[0]

	initMatchLine, err := tools.Map(strings.Split(parts[1], ","), func(t *string) (int, error) {
		return strconv.Atoi(*t)
	})

	tools.AssertNoError(err)

	matchLine := make([]int, 0, len(initMatchLine)*times)

	var patterns []string
	for i := 0; i < times; i++ {
		matchLine = append(matchLine, initMatchLine...)
		patterns = append(patterns, initPattern)
	}

	pattern := strings.Join(patterns, "?")

	return calculateArrangementsOptimal(pattern, matchLine)
}

func extractMatchLine(
	matchLine string,
) (matches []int, hasNoMore bool, nextNum int) {
	matches = make([]int, 0)

	num := 0
	for i := 0; i < len(matchLine); i++ {
		symbol := matchLine[i]

		switch symbol {
		case unknown:
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

func calculateArrangementsOptimal(
	pattern string,
	matchLine []int,
) int {

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

//[???.### 1,1,3] x 5 => [1 1 1 1 1]
//[.??..??...?##. 1,1,3] x 5 => [4 32 256 2048 16384]
//[?#?#?#?#?#?#?#? 1,3,1,6] x 5 => [1 1 1 1 1]
//[????.#...#... 4,1,1] x 5 => [1 2 4 8 16]
//[????.######..#####. 1,6,5] x 5 => [4 20 100 500 2500]
//[?###???????? 3,2,1] x 5 => [10 150 2250 33750 506250]

//[.#???????????#.# 1,5,3,1] 1 => 3 (3)
//[.#???????????#.# 1,5,3,1] 2 => 10 (9)
//[.#???????????#.# 1,5,3,1] 3 => 36 (27)
//[.#???????????#.# 1,5,3,1] 4 => 136 (81)
//[.#???????????#.# 1,5,3,1] 5 => 528 (243)
//[.#???????????#.# 1,5,3,1] 6 => 2080 (729)
//--- PASS: TestAnalysis (0.19s)

//[.??#???.???????? 1,2,1,5] 1 => 10 (0)
//[.??#???.???????? 1,2,1,5] 2 => 314 (0)
//[.??#???.???????? 1,2,1,5] 3 => 10378 (0)
//[.??#???.???????? 1,2,1,5] 4 => 345050 (0)
