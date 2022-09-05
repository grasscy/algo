package ac

type NumMatrix struct {
	pre [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	pre := make([][]int, len(matrix)+1)
	for i := 0; i < len(pre); i++ {
		pre[i] = make([]int, len(matrix[0])+1)
	}
	for i := 1; i < len(pre); i++ {
		for j := 1; j < len(pre[0]); j++ {
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + matrix[i-1][j-1]

		}
	}
	n := NumMatrix{pre}
	return n
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	pre := this.pre
	sum := pre[row2+1][col2+1] - pre[row1][col2+1] - pre[row2+1][col1] + pre[row1][col1]
	return sum
}
