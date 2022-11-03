package ac

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2

	dp := make([]bool, sum+1)
	for i := 1; i <= len(nums); i++ {
		for j := sum; j >= 1; j-- {
			if j >= nums[i-1] {
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
	}

	return dp[sum]
}
