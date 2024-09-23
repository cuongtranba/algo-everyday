package prefixsum

type NumMatrix struct {
	prefixSumRow [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	prefixSumMax := NumMatrix{
		prefixSumRow: make([][]int, 0),
	}
	for row := range matrix {
		sum := 0
		temp := make([]int, len(matrix[row]))
		for col, n := range matrix[row] {
			sum += n
			temp[col] = sum
		}
		prefixSumMax.prefixSumRow = append(prefixSumMax.prefixSumRow, temp)
	}
	return prefixSumMax
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var left, sum int
	for i := row1; i <= row2; i++ {
		if col1 > 0 {
			left = this.prefixSumRow[i][col1-1]
		} else {
			left = 0
		}
		sum += this.prefixSumRow[i][col2] - left
	}
	return sum
}
