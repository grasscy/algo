package jianzhioffer

func minCost(costs [][]int) int {
	dp := initArray(len(costs)+1, 3)
	for i := 1; i < len(dp); i++ {
		dp[i][0] = min(dp[i-1][1], dp[i-1][2]) + costs[i-1][0]
		dp[i][1] = min(dp[i-1][0], dp[i-1][2]) + costs[i-1][1]
		dp[i][2] = min(dp[i-1][1], dp[i-1][0]) + costs[i-1][2]
	}
	ans := min(dp[len(costs)][0], dp[len(costs)][1])
	ans = min(ans, dp[len(costs)][2])
	return ans
}
