package ac

func numRollsToTarget(n int, k int, target int) int {
	dp := initArray(n+1, target+1)
	for i := 1; i <= min(k, target); i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= target; j++ {
			for kk := 1; kk <= k; kk++ {
				if j >= kk {
					dp[i][j] += dp[i-1][j-kk] % (10e8 + 7)
				}
			}
		}
	}
	return dp[n][target] % (10e8 + 7)
}
