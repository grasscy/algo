package ac

func matrixBlockSum(mat [][]int, k int) [][]int {
	pre := make([][]int, len(mat)+1)
	for i := 0; i < len(mat); i++ {
		pre[i] = make([]int, len(mat[0])+1)
	}
	for i := 1; i <= len(mat); i++ {
		for j := 1; j <= len(mat[0]); j++ {
			pre[i][j] = pre[i-1][j] + pre[i][j-1] - pre[i-1][j-1] + mat[i-1][j-1]
		}
	}

	ans := make([][]int, len(mat))
	for i := 0; i < len(mat); i++ {
		ans[i] = make([]int, len(mat[0]))
	}

	for i := 0; i < len(ans); i++ {
		for j := 0; j < len(ans[0]); j++ {
			x1, y1 := max(i-k, 0), max(j-k, 0)
			x2, y2 := min(i+k, len(ans)-1), min(j+k, len(ans[0])-1)
			ans[i][j] = pre[x2+1][y2+1] - pre[x2+1][y1] - pre[x1][y2+1] + pre[x1][y1]
		}
	}
	return ans
}
