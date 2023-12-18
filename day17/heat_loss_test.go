package day17

import (
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTestMath(t *testing.T) {
	lines, err := tools.ReadLinesFileLines("data/test_result.txt")
	assert.Nil(t, err)
	testInput, err := tools.ReadLinesFileLines("data/test.txt")
	assert.Nil(t, err)

	hlm := NewHeatLossMatrix(testInput)

	var positions []Position

	for i, row := range lines {
		for j, symbol := range []byte(row) {
			if !tools.IsAsciiNumber(symbol) {
				positions = append(positions, Position{i, j})
			}
		}
	}

	sum := 0
	for _, pos := range positions {
		sum += hlm.tiles[pos.i][pos.j]
	}

	assert.Equal(t, 102, sum)
	assert.Equal(t, 102, hlm.GetMinimalHeatLoss())
}

func TestNewVisitList(t *testing.T) {
	vl := NewVisitList().Add(Position{0, 0}).Add(Position{0, 1}).Add(Position{0, 2}).Add(Position{0, 3})
	assert.Equal(t, false, vl.SameDirectionLimitReached())

	assert.Equal(t, true, vl.Add(Position{0, 4}).SameDirectionLimitReached())
	assert.Equal(t, false, vl.Add(Position{1, 3}).SameDirectionLimitReached())
}
