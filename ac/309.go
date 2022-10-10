package ac

func maxProfit(prices []int) int {
	dp_0 := 0
	dp_1 := -prices[0]
	dp_pre_0 := 0
	for i := 1; i < len(prices); i++ {
		temp := dp_0
		dp_0 = max(dp_0, dp_1+prices[i])
		dp_1 = max(dp_1, dp_pre_0-prices[i])
		dp_pre_0 = temp

	}
	return dp_0
}
