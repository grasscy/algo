package t200_399

func productExceptSelf(nums []int) []int {
	pres := make([]int, len(nums))
	pre := nums[0]
	pres[0] = 1
	for i := 1; i < len(nums); i++ {
		pres[i] = pre
		pre *= nums[i]
	}

	sufs := make([]int, len(nums))
	suf := nums[len(nums)-1]
	sufs[len(sufs)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		sufs[i] = suf
		suf *= nums[i]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = pres[i] * sufs[i]
	}
	return ans
}
