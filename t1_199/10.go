package t1_199

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] >= 'a' && p[j-1] <= 'z' {
				dp[i][j] = dp[i-1][j-1] && s[i-1] == p[j-1]
			} else if p[j-1] == '*' {
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-1] || dp[i][j-2] || dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
