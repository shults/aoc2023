package day12

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSprings(t *testing.T) {
	tests := []struct {
		line         string
		arrangements int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 16384},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 16},
		{"????.######..#####. 1,6,5", 2500},
		{"?###???????? 3,2,1", 506250},
	}

	for _, test := range tests {
		assert.Equal(t, test.arrangements, CalculateArrangementsPart2(test.line))
	}
}

func Test_calculateArrangementsPart2(t *testing.T) {

	tests := []struct {
		line         string
		arrangements int
	}{
		{"???.### 1,1,3", 1},
		{".??..??...?##. 1,1,3", 4},
		{"?#?#?#?#?#?#?#? 1,3,1,6", 1},
		{"????.#...#... 4,1,1", 1},
		{"????.######..#####. 1,6,5", 4},
		{"?###???????? 3,2,1", 10},
	}

	for _, test := range tests {
		assert.Equal(t, test.arrangements, CalculateArrangements(test.line))
	}
}

func TestAnalysis(t *testing.T) {

	line := ".#???????????#.# 1,5,3,1"
	line = "??.?????????#??? 3,4"
	line = ".??#???.???????? 1,2,1,5"
	line = "?????????#?# 1,1,7"

	for i := 0; i < 5; i++ {
		// CalculateArrangementsPart2parametrized(line, i+1)
		fmt.Printf("[%s] %d => %d (%d)\n", line, i+1, calculateArrangementsPart2(line, i+1), 0)
	}

	//[.#???????????#.# 1,5,3,1] 1 => 3
	//[.#???????????#.# 1,5,3,1] 2 => 10
	//[.#???????????#.# 1,5,3,1] 3 => 36
	//[.#???????????#.# 1,5,3,1] 4 => 136
	//[.#???????????#.# 1,5,3,1] 5 => 528
	//[.#???????????#.# 1,5,3,1] 6 => 2080
	//[.#???????????#.# 1,5,3,1] 7 => 8256
	//[.#???????????#.# 1,5,3,1] 8 => 32896

	//	=== RUN   TestAnalysis
	//[.#???????????#.# 1,5,3,1] 1 => 3
	//[.#???????????#.# 1,5,3,1] 2 => 10
	//[.#???????????#.# 1,5,3,1] 3 => 36
	//[.#???????????#.# 1,5,3,1] 4 => 136
	//[.#???????????#.# 1,5,3,1] 5 => 528
	//[.#???????????#.# 1,5,3,1] 6 => 2080
	//[.#???????????#.# 1,5,3,1] 7 => 8256
	//[.#???????????#.# 1,5,3,1] 8 => 32896
	//[.#???????????#.# 1,5,3,1] 9 => 131328

	//[??.?????????#??? 3,4] 1 => 18 (0)
	//[??.?????????#??? 3,4] 2 => 473 (0)
	//[??.?????????#??? 3,4] 3 => 13893 (0)

}
