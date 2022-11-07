package ac

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	cnt := make([]int, len(nums))
	mm := 1
	for i := 0; i < len(nums); i++ {
		dp[i], cnt[i] = 1, 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
					cnt[i] = cnt[j]
				} else if dp[i] == dp[j]+1 {
					cnt[i] += cnt[j]
				}
			}
		}
		mm = max(mm, dp[i])
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == mm {
			ans += cnt[i]
		}
	}

	return ans

}
