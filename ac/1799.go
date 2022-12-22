package ac

import "math/bits"

func maxScore(nums []int) int {
	ll := len(nums)
	gs := initArray(ll, ll)
	for i := 0; i < ll; i++ {
		for j := 0; j < ll; j++ {
			gs[i][j] = gcd(nums[i], nums[j])
		}
	}
	dp := make([]int, 1<<ll)
	for k := 0; k < 1<<ll; k++ {
		count := bits.OnesCount(uint(k))
		if count%2 == 1 {
			continue
		}
		for i := 0; i < ll; i++ {
			if k>>i&1 == 1 {
				for j := i + 1; j < ll; j++ {
					if k>>j&1 == 1 {
						dp[k] = max(dp[k], dp[k^(1<<i)^(1<<j)]+count/2*gs[i][j])
					}
				}
			}
		}
	}
	return dp[1<<ll-1]
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
