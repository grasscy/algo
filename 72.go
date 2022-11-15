package algo

func minDistance(word1 string, word2 string) int {
	dp := initArray(len(word1)+1, len(word2)+1)
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(word1)][len(word2)]
}
