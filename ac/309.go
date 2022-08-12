package ac

func maxProfit(prices []int) int {
    dp := make([][]int, len(prices))
    for i := 0; i < len(dp); i++ {
        dp[i] = []int{0, 0, 0}
    }
    dp[0][0] = -prices[0] // 持有
    dp[0][1] = 0          // 不持有，冷冻期
    dp[0][2] = 0          // 不持有，不冷冻

    for i := 1; i < len(prices); i++ {
        dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
        dp[i][1] = dp[i-1][0] + prices[i]
        dp[i][2] = max(dp[i-1][1], dp[i-1][2])
    }
    return max(dp[len(dp)-1][1], dp[len(dp)-1][2])

}