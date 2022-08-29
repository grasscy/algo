package ac

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if math.Abs(float64(target)) > math.Abs(float64(sum)) {
		return 0
	}

	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2*sum+1)
	}
	if nums[0] == 0 {
		dp[0][sum] = 2
	} else {
		dp[0][-nums[0]+sum] = 1
		dp[0][nums[0]+sum] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < 2*sum+1; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
			if j+nums[i] < 2*sum+1 {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
		}
	}
	return dp[len(nums)-1][target+sum]
}
