package algo

import "fmt"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	dp := make([]bool, len(nums))
	if nums[1] > nums[0] {
		dp[1] = true
	}
	for i := 2; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			if dp[j] && nums[i] > nums[j] {
				dp[i] = true
				break
			}
		}
	}
	fmt.Println(dp)
	return dp[len(nums)-1]
}
