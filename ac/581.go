package ac

func findUnsortedSubarray(nums []int) int {
	l := 0
	r := -1
	lmax := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] >= lmax {
			lmax = nums[i]
		} else {
			r = i
		}
	}
	rmin := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] <= rmin {
			rmin = nums[i]
		} else {
			l = i
		}
	}
	return r - l + 1
}
