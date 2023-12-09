package day09_test

import (
	"aoc2023/day09"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSequence(t *testing.T) {
	s, err := day09.ParseSequence("10  13  16  21  30")
	assert.Nil(t, err)
	assert.Equal(t, 30, s.GetNextByShift(0))
	assert.Equal(t, 45, s.GetNextByShift(1))
	assert.Equal(t, 68, s.GetNextByShift(2))
	assert.Equal(t, 101, s.GetNextByShift(3))
}

func TestParseSequence(t *testing.T) {
	s, err := day09.ParseSequence("10  13  16  21  30  45")
	assert.Nil(t, err)
	assert.Equal(t, 45, s.GetNextByShift(0))
	assert.Equal(t, 68, s.GetNextByShift(1))
	assert.Equal(t, 101, s.GetNextByShift(2))

	assert.Equal(t, 10, s.GetPrevByShift(0)) // todo: add to upper case
	assert.Equal(t, 5, s.GetPrevByShift(1))  // todo: add to upper case
}
