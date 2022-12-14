package t200_399

import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	small, mid := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v <= small {
			small = v
		} else if v <= mid {
			mid = v
		} else if v > mid {
			return true
		}
	}

	return false
}
