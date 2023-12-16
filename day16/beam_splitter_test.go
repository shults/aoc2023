package day16

import (
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBeamSplitter(t *testing.T) {
	lines, err := tools.ReadLinesFileLines("test.txt")
	assert.Nil(t, err)

	bs := NewBeamSplitter(lines)

	assert.Equal(t, 46, bs.CalculateEnergizedTiles(false))
}
