package tools

func MatrixInsertRow[T any](matrix *[][]T, rowNumber int, row []T) {
	mat := *matrix

	rowLen := len(mat[0])

	mat = append(mat, make([]T, rowLen))

	for i := len(mat) - 1; i > rowNumber; i-- {
		copy(mat[i], mat[i-1])
	}

	AssertTrue(len(mat[rowNumber]) == len(row), "expected row sith same size")

	for j, _ := range (mat)[rowNumber] {
		mat[rowNumber][j] = row[j]
	}

	*matrix = mat
}

func MatrixInsertCol[T any](matrix *[][]T, colNum int, col []T) {
	mat := *matrix
	dim := MatrixSize(mat)

	for i, _ := range mat {
		var val T
		mat[i] = append(mat[i], val)

		for j := dim.Cols; j > colNum; j-- {
			mat[i][j] = mat[i][j-1]
		}
	}

	for i := 0; i < dim.Rows; i++ {
		mat[i][colNum] = col[i]
	}

	*matrix = mat
}

type Dimensions struct {
	Rows, Cols int
}

func (d *Dimensions) Contains(i, j int) bool {
	return i >= 0 && j >= 0 && i < d.Rows && j < d.Cols
}

func MatrixSize[T any](matrix [][]T) (dim Dimensions) {
	dim.Rows = len(matrix)

	for i, row := range matrix {
		if i == 0 {
			dim.Cols = len(row)
			continue
		}

		if dim.Cols != len(row) {
			panic("square matrix expected")
		}
	}

	return
}
