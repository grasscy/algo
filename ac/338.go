package ac

func countBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	cur := 1

	for i := 2; i <= n; i++ {
		if i == cur*2 {
			dp[i] = 1
			cur *= 2
		} else {
			dp[i] = dp[cur] + dp[i-cur]
		}
	}
	return dp
}
