package day12

import (
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

	//assert.Equal(t, 1, calculateArrangementsOptimal("???.###", []int{1, 1, 3}))
	//assert.Equal(t, 4, calculateArrangementsOptimal(".??..??...?##.", []int{1, 1, 3}))
}
