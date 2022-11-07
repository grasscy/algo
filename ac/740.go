package ac

func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	all := make([]int, 100001)
	for _, v := range nums {
		all[v]++
	}
	dp := make([]int, 100001)
	dp[1] = all[1]
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+i*all[i])
	}
	return dp[len(dp)-1]
}
