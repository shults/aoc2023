package day01

import (
	"strings"
)

type wordsMap = map[string]int

var wordsToDigit wordsMap
var invertedMap wordsMap

func init() {
	wordsToDigit = wordsMap{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	invertedMap = makeInvertedMap(wordsToDigit)
}

func makeInvertedMap(input wordsMap) wordsMap {
	res := make(wordsMap, len(input))

	for key, val := range input {
		res[reverse(key)] = val
	}

	return res
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func isAsciiDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func asciiDigitToInt(ch byte) int {
	return int(ch - '0')
}

func GetTwoDigitsNumber(line string) int {
	return parseFirstDigit(line)*10 + parseLastDigit(line)
}

func parseFirstDigit(line string) int {
	return parseDigit(line, wordsToDigit)
}

func parseLastDigit(line string) int {
	return parseDigit(reverse(line), invertedMap)
}

func parseDigit(
	line string,
	wm wordsMap,
) int {
	i := 0
	var ch byte = 0

	for ; i < len(line); i++ {
		ch = (line)[i]

		if !isAsciiDigit(ch) {
			continue
		}

		break
	}

	firstWordIndex := -1
	var firstWordValue int

	for sDigit, digitVal := range wm {
		index := strings.Index(line[0:i], sDigit)

		if index == -1 {
			continue
		}

		if firstWordIndex == -1 || index < firstWordIndex {
			firstWordIndex = index
			firstWordValue = digitVal
		}
	}

	if firstWordIndex == -1 {
		return asciiDigitToInt(ch)
	} else {
		return firstWordValue
	}
}
