package ac

func longestSubsequence(arr []int, difference int) int {
	dp := map[int]int{}
	ans := 1
	for _, v := range arr {
		dp[v] = dp[v-difference] + 1
		if ans < dp[v] {
			ans = dp[v]
		}
	}
	return ans
}
