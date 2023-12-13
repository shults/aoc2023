package day13_test

import (
	"aoc2023/day13"
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	data, err := tools.ReadLinesFileLines("test.txt")
	assert.Nil(t, err)
	matrices := day13.NewMatrices(data)
	assert.Equal(t, 2, len(matrices))
	assert.Equal(t, 405, matrices.CalculatePart1())
}

func TestMyCase(t *testing.T) {
	data, err := tools.ReadLinesFileLines("my.txt")
	assert.Nil(t, err)
	matrices := day13.NewMatrices(data)
	assert.Equal(t, 2, len(matrices))
	assert.Equal(t, 405, matrices.CalculatePart1())
}
