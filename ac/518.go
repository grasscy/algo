package ac

func change(amount int, coins []int) int {
	dp := initArray(amount+1, len(coins)+1)
	for j := 0; j < len(coins)+1; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < amount+1; i++ {
		for j := 1; j < len(coins)+1; j++ {
			if i-coins[j-1] >= 0 {
				dp[i][j] += dp[i-coins[j-1]][j]
			}
			dp[i][j] += dp[i][j-1]
		}
	}
	return dp[amount][len(coins)]

}
