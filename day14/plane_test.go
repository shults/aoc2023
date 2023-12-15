package day14_test

import (
	"aoc2023/day14"
	"aoc2023/tools"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaneTestSimple(t *testing.T) {
	data, err := tools.ReadLinesFileLines("test.txt")
	assert.Nil(t, err)
	plane := day14.NewPlane(data)
	assert.Equal(t, 136, plane.CalculatePart1(false))
}

func TestPlaneTestRotation(t *testing.T) {
	t.Skip()

	plane := day14.NewPlane([]string{
		"123",
		"456",
		"789",
	})

	fmt.Printf("%s\n", plane.String())

	plane.RotateRight()
	fmt.Printf("%s\n", plane.String())

	plane.RotateRight()
	fmt.Printf("%s\n", plane.String())

	plane.RotateRight()
	fmt.Printf("%s\n", plane.String())

	plane.RotateRight()
	fmt.Printf("%s\n", plane.String())
}
