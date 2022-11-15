package ac

func maxSubArray(nums []int) int {
	pre := nums[0]
	mm := pre
	for i := 1; i < len(nums); i++ {
		pre = max(pre+nums[i], nums[i])
		if mm < pre {
			mm = pre
		}
	}
	return mm
}
