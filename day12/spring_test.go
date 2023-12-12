package day12_test

import (
	"aoc2023/day12"
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
		assert.Equal(t, test.arrangements, day12.CalculateArrangements(test.line))
	}

}
