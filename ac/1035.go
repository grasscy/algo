package ac

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := initArray(len(nums1)+1, len(nums2)+1)
	mm := 0
	for i := 1; i <= len(nums1); i++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if mm < dp[i][j] {
				mm = dp[i][j]
			}
		}
	}
	return mm
}
