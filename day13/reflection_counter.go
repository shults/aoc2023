package day13

import (
	"aoc2023/tools"
)

func createIdGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

// part1=27742
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

func (m Matrices) CalculatePart1() int {
	res := 0

	for _, matrix := range m {
		res += matrix.CalculatePart1()
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

func (m *Matrix) CalculatePart1() int {
	row := m.getRowReflectionLine()
	col := m.getColReflectionLine()

	if row == 0 {
		return col * 100
	}

	return row
}

func (m *Matrix) getRowReflectionLine() int {
	rowReflectionLine := 0

	for j := 0; j < m.dim.Cols-1; j++ {
		rowReflectionLine = tools.Max(rowReflectionLine, m.findColReflection(j, j+1))
	}

	return rowReflectionLine
}

func (m *Matrix) findColReflection(a, b int) int {
	matched := a + 1

	for {
		if m.compareCols(a, b) {
			if a == 0 || b == m.dim.Cols-1 {
				return matched
			}

			a, b = a-1, b+1
		} else {
			return 0
		}
	}
}

func (m *Matrix) compareCols(a, b int) bool {
	for i := 0; i < m.dim.Rows; i++ {
		if m.data[i][a] != m.data[i][b] {
			return false
		}
	}

	return true
}

func (m *Matrix) getColReflectionLine() int {
	colReflection := 0

	for i := 0; i < m.dim.Rows-1; i++ {
		colReflection = tools.Max(colReflection, m.findRowReflection(i, i+1))
	}

	return colReflection
}

func (m *Matrix) findRowReflection(a, b int) int {
	matched := a + 1

	for {
		if m.compareRows(a, b) {
			if a == 0 || b == m.dim.Rows-1 {
				return matched
			}

			a, b = a-1, b+1
		} else {
			return 0
		}
	}
}

func (m *Matrix) compareRows(a, b int) bool {
	for j := 0; j < m.dim.Cols; j++ {
		if m.data[a][j] != m.data[b][j] {
			return false
		}
	}

	return true
}
