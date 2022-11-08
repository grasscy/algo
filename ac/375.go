package ac

import "math"

var memo [][]int

func getMoneyAmount(n int) int {
	memo = initArray(n+1, n+1)
	return dp(1, n)
}

func dp(lo, hi int) int {
	if lo >= hi {
		return 0
	}
	if lo == hi-1 {
		return lo
	}
	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}
	ans := math.MaxInt32
	for i := lo; i <= hi; i++ {
		ans = min(ans, max(dp(lo, i-1), dp(i+1, hi))+i)
	}
	memo[lo][hi] = ans
	return ans
}
