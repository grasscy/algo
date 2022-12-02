package t1_199

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := math.MaxInt32

	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if math.Abs(float64(sum-target)) < math.Abs(float64(ans-target)) {
				ans = sum
			}
			if sum >= target {
				k--
			} else {
				j++
			}
		}
	}
	return ans

}
