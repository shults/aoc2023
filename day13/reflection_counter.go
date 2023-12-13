package day13

import (
	"aoc2023/tools"
	"fmt"
)

func createIdGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// part1=27742
// part2=32728
func NewMatrices(data []string) Matrices {
	var matrices Matrices
	var matrixData []string
	idGenerator := createIdGenerator()

	for _, line := range data {
		if len(line) == 0 {
			matrices = append(matrices, NewMatrix(matrixData, idGenerator()))
			matrixData = nil
			continue
		}

		matrixData = append(matrixData, line)
	}

	if matrixData != nil {
		matrices = append(matrices, NewMatrix(matrixData, idGenerator()))
	}

	return matrices
}

type Matrices []Matrix

func (m Matrices) CalculatePart1(verbose bool) int {
	res := 0

	for _, matrix := range m {
		res += matrix.CalculatePart1(verbose)
	}

	return res
}

func (m Matrices) CalculatePart2(verbose bool) int {
	res := 0

	for _, matrix := range m {
		res += matrix.CalculatePart2(verbose)
	}

	return res
}

type Matrix struct {
	data [][]byte
	dim  tools.Dimensions
	id   int
}

func NewMatrix(data []string, id int) Matrix {
	var bytes [][]byte

	for _, line := range data {
		bytes = append(bytes, []byte(line))
	}

	return Matrix{
		data: bytes,
		dim:  tools.MatrixSize(bytes),
		id:   id,
	}
}

func (m *Matrix) CalculatePart1(verbose bool) int {
	row := m.getReflectionLine(0)
	m.swapRowsWithColumns()
	col := m.getReflectionLine(0)
	m.swapRowsWithColumns()

	if verbose {
		fmt.Printf("matrix[%d] col=%d row=%d \n", m.id, col, row)
	}

	return col + row*100
}

func (m *Matrix) CalculatePart2(verbose bool) int {
	row := m.getReflectionLine(1)
	m.swapRowsWithColumns()
	col := m.getReflectionLine(1)
	m.swapRowsWithColumns()

	if verbose {
		fmt.Printf("matrix[%d] col=%d row=%d \n", m.id, col, row)
	}

	return col + row*100
}

func (m *Matrix) swapRowsWithColumns() {
	data := make([][]byte, m.dim.Cols)

	for j := 0; j < m.dim.Cols; j++ {
		data[j] = make([]byte, m.dim.Rows)
	}

	for i := 0; i < m.dim.Rows; i++ {
		for j := 0; j < m.dim.Cols; j++ {
			data[j][i] = m.data[i][j]
		}
	}

	m.dim.Cols, m.dim.Rows = m.dim.Rows, m.dim.Cols
	m.data = data
}

func (m *Matrix) getReflectionLine(allowedMismatches int) int {

	for i := 0; i < m.dim.Rows-1; i++ {
		reflectionRow := m.findReflection(i, i+1, allowedMismatches)

		if reflectionRow != 0 {
			return reflectionRow
		}
	}

	return 0
}

func (m *Matrix) findReflection(a, b, allowedMismatches int) int {
	matched := a + 1
	diffsSum := 0

	for {
		diffs := m.countDiffs(a, b)
		diffsSum += diffs

		if diffsSum <= allowedMismatches {
			if a == 0 || b == m.dim.Rows-1 {
				if diffsSum == allowedMismatches {
					return matched
				} else {
					return 0
				}
			}
			a, b = a-1, b+1
		} else {
			return 0
		}
	}
}

func (m *Matrix) countDiffs(a, b int) int {
	diffs := 0

	for j := 0; j < m.dim.Cols; j++ {
		if m.data[a][j] != m.data[b][j] {
			diffs++
		}
	}

	return diffs
}
