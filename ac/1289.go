package ac

import "math"

func minFallingPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := initArray(m, n)
	for j := 0; j < n; j++ {
		dp[0][j] = grid[0][j]
	}
	// 记录上一行的最小值下标和次小值下标
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			mmin := math.MaxInt32
			for k := 0; k < n; k++ {
				if j != k && dp[i-1][k] < mmin {
					mmin = dp[i-1][k]
				}
			}
			dp[i][j] = mmin + grid[i][j]
		}
	}
	ans := math.MaxInt32
	for j := 0; j < n; j++ {
		if ans > dp[m-1][j] {
			ans = dp[m-1][j]
		}
	}
	return ans
}
