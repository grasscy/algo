package ac

func change(amount int, coins []int) int {
	dp := initArray(amount+1, len(coins)+1)
	for i := 0; i <= len(coins); i++ {
		dp[0][i] = 1
	}

	for i := 1; i <= amount; i++ {
		for j := 1; j <= len(coins); j++ {
			if coins[j-1] > i {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-coins[j-1]][j]
			}
		}
	}
	return dp[amount][len(coins)]
}
