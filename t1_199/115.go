package t1_199

func numDistinct(s string, t string) int {
	dp := initArray(len(s)+1, len(t)+1)
	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(t); j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(s)][len(t)]
}
