package t1_199

import "math"

func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}

	return ans
}
