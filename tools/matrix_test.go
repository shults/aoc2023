package tools_test

import (
	"aoc2023/tools"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatrixInsertRow(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}

	tools.MatrixInsertRow(&matrix, 1, []int{101, 102})

	assert.Equal(t, [][]int{
		{1, 2},
		{101, 102},
		{3, 4},
	}, matrix)

	assert.Panics(t, func() {
		tools.MatrixInsertRow(&matrix, 1, []int{101, 102, 300})
	})
}

func TestMatrixInsertCol(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	tools.MatrixInsertCol(&matrix, 2, []int{100, 100})

	assert.Equal(t, [][]int{
		{1, 2, 100, 3},
		{4, 5, 100, 6},
	}, matrix)
}
