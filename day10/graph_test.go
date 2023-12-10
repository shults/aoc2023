package day10

import (
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSymbol(t *testing.T) {
	assert.Panics(t, func() {
		NewTileSymbol('a')
	})
}

func TestGrafConstruction(t *testing.T) {
	g := NewGraph([]string{
		"......",
		".S--7.",
		".|..|.",
		".|--|.",
		".L--J.",
		"......",
	})

	assert.NotNil(t, g.startTile)
	assert.Equal(t, 6, g.CalculatePart1())
	assert.Equal(t, 4, g.CalculatePart2())
}

func TestGrafConstructionPart1MoreComplexGraph(t *testing.T) {
	g := NewGraph([]string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	})

	assert.NotNil(t, g.startTile)
	assert.Equal(t, 8, g.CalculatePart1())
}

func TestGrafConstructionPart1MoreComplexGraphPart2(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{"test_part2_simple.txt", 8},
		{"test_part2_last.txt", 10},
		{"my.txt", 579},
	}

	for _, test := range tests {
		data, err := tools.ReadFile(test.file)
		assert.Nil(t, err)

		g := NewGraph(data)

		assert.NotNil(t, g.startTile)
		assert.Equal(t, test.expected, g.CalculatePart2())
	}
}
