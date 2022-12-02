package t1_199

import "strconv"

func numDecodings(s string) int {
	dp := make([]int, len(s))
	if s[0] == '0' {
		return 0
	} else {
		dp[0] = 1
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if s[i-1] == '0' {
				dp[i] = 0
				break
			}
			atoi, _ := strconv.Atoi(s[i-1 : i+1])
			if atoi > 26 {
				dp[i] = 0
				break
			}
			if i-2 >= 0 {
				dp[i] = dp[i-2]
			} else {
				dp[i] = 1
			}
		} else {
			if s[i-1] == '0' {
				dp[i] = dp[i-1]
			} else {
				atoi, _ := strconv.Atoi(s[i-1 : i+1])
				if atoi <= 26 {
					if s[i-1] == '0' {

					}
					if i-2 >= 0 {
						dp[i] = dp[i-1] + dp[i-2]
					} else {
						dp[i] = 2
					}
				} else {
					dp[i] = dp[i-1]
				}
			}
		}
	}
	return dp[len(dp)-1]
}
