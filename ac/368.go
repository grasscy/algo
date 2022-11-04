package ac

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = []int{nums[i]}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(dp[j])+1 > len(dp[i]) {
				dp[i] = append([]int{}, dp[j]...)
				dp[i] = append(dp[i], nums[i])
			}
		}
	}
	ans := 0
	max := 0
	for i := 0; i < len(dp); i++ {
		if max < len(dp[i]) {
			ans = i
			max = len(dp[i])
		}
	}
	return dp[ans]
}
