package day10

import (
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
		".....",
		".S-7.",
		".|.|.",
		".L-J.",
		".....",
	})

	g.CalculatePart2()

	assert.NotNil(t, g.startTile)
	assert.Equal(t, 4, g.CalculatePart1())
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
