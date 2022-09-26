package ac

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = 999
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if j < 1 {
				dp[i+1][j] = min(dp[i][j], dp[i][j+1])
			} else if j == len(matrix[0])-1 {
				dp[i+1][j] = min(dp[i][j], dp[i][j-1])
			} else {
				dp[i+1][j] = min(dp[i][j], dp[i][j-1])
				dp[i+1][j] = min(dp[i+1][j], dp[i][j+1])
			}
			dp[i+1][j] += matrix[i+1][j]
		}
	}
	ans := 99999
	for j := 0; j < len(matrix[0]); j++ {
		ans = min(ans, dp[len(matrix)-1][j])
	}
	return ans
}
