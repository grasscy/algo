package pass

func rob(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = dp[i-1]
		if dp[i-2]+nums[i-1] > dp[i] {
			dp[i] = dp[i-2] + nums[i-1]
		}
	}
	return dp[len(nums)]
}
