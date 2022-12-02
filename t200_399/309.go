package t200_399

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	if len(prices) == 2 {
		return max(prices[1]-prices[0], 0)
	}

	dp := initArray(len(prices), 3) // 当天结束：0持有，1不持有，2冷冻
	dp[0][0] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][2], dp[i-1][1])
		dp[i][2] = dp[i-1][0] + prices[i]
	}

	return max(dp[len(prices)-1][1], dp[len(prices)-1][2])
}
