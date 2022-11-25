package jianzhioffer

type NumMatrix struct {
	pre [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	p := make([][]int, len(matrix)+1)
	for i := 0; i < len(p); i++ {
		p[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			p[i][j] = p[i-1][j] + p[i][j-1] - p[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	nm := NumMatrix{pre: p}
	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.pre[row2+1][col2+1] - this.pre[row1][col2+1] - this.pre[row2+1][col1] + this.pre[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
