package ac

func maxProfit(prices []int) int {
	pre0 := 0
	pre1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		pre0 = max(pre0, pre1+prices[i])
		pre1 = max(pre1, -prices[i])
	}
	return pre0
}
