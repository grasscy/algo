package ac

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)
	left, right := 0, nums[len(nums)-1]
	for left <= right {
		mid := (left + right) / 2
		d := f(nums, mid)
		if d == mid {
			return mid
		} else if d < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func f(nums []int, x int) int {
	i := 0
	for ; i < len(nums) && nums[i] < x; i++ {

	}
	return len(nums) - i
}
