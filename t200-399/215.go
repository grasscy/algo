package t200_399

func findKthLargest(nums []int, k int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		kk := q(nums, lo, hi)
		if kk == len(nums)-k {
			return nums[kk]
		} else if kk > len(nums)-k {
			hi = kk - 1
		} else {
			lo = kk + 1
		}
	}
	return -1
}

func q(nums []int, lo, hi int) int {
	p := lo
	for i := lo + 1; i <= hi; i++ {
		if nums[lo] > nums[i] {
			p++
			nums[i], nums[p] = nums[p], nums[i]
		}
	}
	nums[lo], nums[p] = nums[p], nums[lo]

	return p
}
