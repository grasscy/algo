package ac

func findSubstringInWraproundString(p string) int {
	count := map[uint8]int{}
	dp := make([]int, len(p))
	dp[0] = 1
	count[p[0]] = 1
	for i := 1; i < len(p); i++ {
		if p[i]-p[i-1] == 1 || (p[i] == 'a' && p[i-1] == 'z') {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
		count[p[i]] = max(count[p[i]], dp[i])
	}
	ans := 0
	for _, v := range count {
		ans += v
	}
	return ans
}
