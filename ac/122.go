package ac

func maxProfit(prices []int) int {
	ans := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			ans += prices[i+1] - prices[i]
		}
	}
	return ans
}
