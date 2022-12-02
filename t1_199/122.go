package t1_199

func maxProfit(prices []int) int {
	pre0 := 0
	pre1 := -prices[0]
	for i := 1; i < len(prices); i++ {
		temp := pre0
		pre0 = max(pre0, pre1+prices[i])
		pre1 = max(pre1, temp-prices[i])
	}
	return pre0
}
