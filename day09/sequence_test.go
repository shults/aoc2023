package day09_test

import (
	"aoc2023/day09"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSequence(t *testing.T) {
	s := day09.NewSequence(30, []int{21, 5, 2, 2})
	assert.Equal(t, 30, s.GetByShift(0))
	assert.Equal(t, 45, s.GetByShift(1))
	assert.Equal(t, 68, s.GetByShift(2))
	assert.Equal(t, 101, s.GetByShift(3))
}

func TestParseSequence(t *testing.T) {
	s, _ := day09.ParseSequence("10  13  16  21  30  45")
	assert.Equal(t, 45, s.GetByShift(0))
	assert.Equal(t, 68, s.GetByShift(1))
	assert.Equal(t, 101, s.GetByShift(2))
}
