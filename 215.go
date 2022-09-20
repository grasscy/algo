package algo

import "math/rand"

func findKthLargest(nums []int, k int) int {

	lo, hi := 0, len(nums)-1
	for lo <= hi {
		p := qs(nums, lo, hi)
		if p == len(nums)-k {
			return nums[p]
		} else if p < len(nums)-k {
			lo = p + 1
		} else {
			hi = p - 1
		}
	}
	return -1
}

func qs(nums []int, lo int, hi int) int {
	p := rand.Intn(hi-lo+1) + lo
	nums[p], nums[lo] = nums[lo], nums[p]

	pivot := nums[lo]
	lt := lo
	for i := lo + 1; i <= hi; i++ {
		if nums[i] < pivot {
			lt++
			nums[i], nums[lt] = nums[lt], nums[i]
		}
	}
	nums[lt], nums[lo] = nums[lo], nums[lt]
	return lt
}
