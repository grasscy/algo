package ac

func maxTurbulenceSize(arr []int) int {
	dp := initArray(len(arr), 2)
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			if i%2 == 1 {
				dp[i][0] = dp[i-1][0] + 1
				dp[i][1] = 1
			} else {
				dp[i][1] = dp[i-1][1] + 1
				dp[i][0] = 1
			}
		} else if arr[i] < arr[i-1] {
			if i%2 == 1 {
				dp[i][1] = dp[i-1][1] + 1
				dp[i][0] = 1
			} else {
				dp[i][0] = dp[i-1][0] + 1
				dp[i][1] = 1
			}
		} else {
			dp[i][0] = 1
			dp[i][1] = 1
		}
	}

	ans := 1
	for _, v := range dp {
		i := max(v[0], v[1])
		ans = max(ans, i)
	}

	return ans
}
