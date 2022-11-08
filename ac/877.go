package ac

func stoneGame(piles []int) bool {
	dp := make([][][]int, len(piles))
	for i := 0; i < len(dp); i++ {
		dp[i] = initArray(len(piles), 2)
		dp[i][i][0] = piles[i]
	}
	for i := len(piles) - 2; i >= 0; i-- {
		for j := i + 1; j < len(piles); j++ {
			left := piles[i] + dp[i+1][j][1]
			right := piles[j] + dp[i][j-1][1]
			dp[i][j][0] = max(left, right)
			if left > right {
				dp[i][j][1] = dp[i+1][j][0]
			} else {
				dp[i][j][1] = dp[i][j-1][0]
			}
		}
	}
	ints := dp[0][len(piles)-1]
	return ints[0]-ints[1] > 0
}
