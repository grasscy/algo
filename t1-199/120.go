package pass

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		for j := 0; j < i+1; j++ {
			v := triangle[i][j]
			dp[i][j] = math.MaxInt32
			if j != 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+v)
			}
			if j != i {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+v)
			}
		}
	}
	ret := math.MaxInt32
	for j := 0; j < n; j++ {
		ret = min(ret, dp[n-1][j])
	}
	return ret
}
