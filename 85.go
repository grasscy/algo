package algo

import "fmt"

type N struct {
	len  int
	high int
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]N, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]N, n)
	}
	for i := 0; i < m; i++ {
		if matrix[0][i] == 1 {
			dp[0][i].len = 1
			dp[0][i].high = 1
		}
		if matrix[i][0] == 1 {
			dp[i][0].len = 1
			dp[i][0].high = 1
		}
	}
	mmax := 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				continue
			}
			dp[i][j].len = min(min(dp[i][j-1].len, dp[i-1][j-1].len), dp[i-1][j].len) + 1
			dp[i][j].high = min(min(dp[i-1][j].high, dp[i-1][j-1].high), dp[i][j-1].high) + 1
			if mmax < dp[i][j].len*dp[i][j].high {
				mmax = dp[i][j].len * dp[i][j].high
			}
		}
	}
	fmt.Println(dp)
	return mmax
}
