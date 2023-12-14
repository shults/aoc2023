package day14_test

import (
	"aoc2023/day14"
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaneTestSimple(t *testing.T) {
	data, err := tools.ReadLinesFileLines("test.txt")
	assert.Nil(t, err)
	plane := day14.NewPlane(data)
	assert.Equal(t, 136, plane.CalculatePart1(false))
}
