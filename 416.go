package algo

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
	dp[0] = true
	for i := 1; i <= sum; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] = dp[i-nums[j]]
				if dp[i] {
					break
				}
			}
		}
	}
	return dp[sum]
}
