package ac

func maxProfit(prices []int, fee int) int {
	pre0 := 0
	pre1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		temp := pre0
		pre0 = max(pre0, pre1+prices[i]-fee)
		pre1 = max(pre1, temp-prices[i])
	}
	return pre0
}
