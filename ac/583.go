package ac

var memo [][]int

func minDistance(word1 string, word2 string) int {
	memo = make([][]int, len(word1))
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, len(word2))
	}

	fill2(memo, -1)

	subsequence := longestCommonSubsequence(word1, word2)
	ans := len(word1) - subsequence + len(word2) - subsequence
	return ans
}

func longestCommonSubsequence(word1 string, word2 string) int {
	return dp(word1, word2, 0, 0)
}

func dp(word1 string, word2 string, m int, n int) int {
	if len(word1) == m || len(word2) == n {
		return 0
	}
	if memo[m][n] != -1 {
		return memo[m][n]
	}

	if word1[m] == word2[n] {
		memo[m][n] = 1 + dp(word1, word2, m+1, n+1)
	} else {
		memo[m][n] = max(dp(word1, word2, m+1, n), dp(word1, word2, m, n+1))
	}
	return memo[m][n]
}
