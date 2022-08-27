package ac

func canPartition(nums []int) bool {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	if nums[0] <= target {
		dp[nums[0]] = true
	}

	for i := 1; i < len(nums); i++ {
		for j := target; nums[i] <= j; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}
