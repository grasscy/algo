package to

func longestPalindrome(s string) string {
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
    }
    for i := 0; i < len(s); i++ {
        dp[i][i] = true
    }

    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s); j++ {

        }
    }
}
