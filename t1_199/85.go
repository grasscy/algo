package t1_199

type N struct {
	len  int
	high int
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]N, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]N, n)
	}
	mmax := 0
	//for i := 0; i < m; i++ {
	//	if matrix[i][0] == '1' {
	//		dp[i][0].len = 1
	//		dp[i][0].high = 1
	//		mmax = 1
	//	}
	//}
	//for i := 0; i < n; i++ {
	//	if matrix[0][i] == '1' {
	//		dp[0][i].len = 1
	//		dp[0][i].high = 1
	//		mmax = 1
	//	}
	//}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i > 0 {
				dp[i][j].high = dp[i-1][j].high + 1
			} else {
				dp[i][j].high = 1
			}
			//dp[i][j].len = dp[i][j-1].len + 1

			l := dp[i][j].high
			s := 0
			for k := j; k >= 0 && matrix[i][k] != '0'; k-- {
				l = min(dp[i][k].high, l)
				s = l * (j - k + 1)
				if s > mmax {
					mmax = s
				}
			}

		}
	}
	return mmax
}
