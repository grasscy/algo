package ac

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	pre := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(pre+nums[i], nums[i])
		pre = dp[i]
	}

	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}
