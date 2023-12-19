package day18

import (
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_newTrenchDigger_test_case(t *testing.T) {
	lines, err := tools.ReadLinesFileLines("data/test.txt")
	assert.Nil(t, err)

	td := newTrenchDigger(lines)
	assert.Equal(t, 62, td.square(true))
}

func Test_newTrenchDigger_my_case(t *testing.T) {
	lines, err := tools.ReadLinesFileLines("data/my.txt")
	assert.Nil(t, err)

	td := newTrenchDigger(lines)
	assert.Equal(t, 62, td.square(true))
}
