package day11_test

import (
	"aoc2023/day11"
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniverse(t *testing.T) {
	for _, test := range []struct {
		file  string
		part1 int
	}{
		{"test.txt", 374},
	} {
		lines, err := tools.ReadLinesFileLines(test.file)
		assert.Nil(t, err)
		universe := day11.NewUniverse(lines)
		assert.Equal(t, test.part1, universe.Part1())
	}
}
