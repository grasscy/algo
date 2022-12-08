package t200_399

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	ans := math.MaxInt32
	sum := 0
	for ; r < len(nums); r++ {
		sum += nums[r]
		for sum-nums[l] >= target {
			sum -= nums[l]
			l++
		}
		if sum >= target {
			ans = min(ans, r-l+1)
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
