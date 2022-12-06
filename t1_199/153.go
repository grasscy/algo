package t1_199

import "math"

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	ans := math.MaxInt32
	for l <= r {
		mid := (l + r) / 2
		//  5 6 1 2 3 4
		//  4 5 6 1 2 3
		ans = min(ans, nums[l])
		ans = min(ans, nums[mid])
		if nums[l] > nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}
