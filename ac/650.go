package ac

func minSteps(n int) int {
	dp := make([]int, n+1)
	dp[1] = 0
	for i := 2; i <= n; i++ {
		for j := i - 1; j >= 1; j-- {
			if i%j == 0 {
				dp[i] = dp[j] + i/j
				break
			}
		}
	}
	return dp[n]
}
