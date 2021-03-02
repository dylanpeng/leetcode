package result

type NumMatrix struct {
	Sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	for i := 0; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	return NumMatrix{
		Sums: matrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	result := 0
	for i := row1; i <= row2; i++ {
		if col1 == 0 {
			result += this.Sums[i][col2]
		} else {
			result += this.Sums[i][col2] - this.Sums[i][col1-1]
		}
	}

	return result
}
