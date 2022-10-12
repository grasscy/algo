package ac

func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		dp[i][i][0] = piles[i]
		dp[i][i][1] = 0
	}

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			left := piles[i] + dp[i+1][j][1]
			right := piles[j] + dp[i][j-1][1]
			if left > right {
				dp[i][j][0] = left
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][0] = right
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}
	return dp[0][n-1][0]-dp[0][n-1][1] > 0
}
