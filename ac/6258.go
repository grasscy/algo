package ac

import "sort"

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	fill1(dp, 1)

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j]*nums[j] > nums[i] {
				break
			}
			if nums[j]*nums[j] == nums[i] {
				dp[i] = dp[j] + 1
				break
			}
		}
	}

	ans := -1
	for _, v := range dp {
		if v > 1 && ans < v {
			ans = v
		}
	}
	return ans

}
