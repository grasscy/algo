package ac

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(rr(nums, 0, len(nums)-1), rr(nums, 1, len(nums)))
}

func rr(nums []int, s, e int) int {
	dp_pre2 := nums[s]
	dp_pre1 := max(dp_pre2, nums[s+1])
	for i := s + 2; i < e; i++ {
		temp := dp_pre1
		dp_pre1 = max(dp_pre2+nums[i], dp_pre1)
		dp_pre2 = temp
	}
	return dp_pre1
}
