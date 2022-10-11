package ac

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp_pre2 := nums[0]
	dp_pre1 := max(dp_pre2, nums[1])
	for i := 2; i < len(nums); i++ {
		temp := dp_pre1
		dp_pre1 = max(dp_pre2+nums[i], dp_pre1)
		dp_pre2 = temp
	}
	return dp_pre1
}
