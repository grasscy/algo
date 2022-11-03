package ac

func findMaxForm(strs []string, m int, n int) int {
	//dp := make([][][]int, len(strs)+1)
	//for i := 0; i < len(dp); i++ {
	//    dp[i] = initArray(m+1, n+1)
	//}
	dp := initArray(m+1, n+1)
	for i := 1; i <= len(strs); i++ {
		n0, n1 := c(strs[i-1])
		for j := m; j >= 0; j-- {
			for k := n; k >= 0; k-- {
				if n0 > j || n1 > k {
					dp[j][k] = dp[j][k]
				} else {
					dp[j][k] = max(dp[j][k], dp[j-n0][k-n1]+1)
				}
			}
		}
	}
	return dp[m][n]
}

func c(s string) (n0 int, n1 int) {
	for _, v := range s {
		if v == '0' {
			n0++
		} else {
			n1++
		}
	}
	return n0, n1
}
