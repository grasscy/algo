package pass

func maxProfit(prices []int) int {
	dp := make([][][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 2+1)
		for j := 0; j <= 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for k := 1; k <= 2; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][2][0]
}
