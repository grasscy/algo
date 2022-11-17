package ac

func maxProfit(prices []int) int {
	mmin := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-mmin > ans {
			ans = prices[i] - mmin
		}
		if prices[i] < mmin {
			mmin = prices[i]
		}
	}
	return ans
}
