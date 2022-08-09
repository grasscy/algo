package algo

// opt
func longestPalindrome(s string) string {
    dp := make([][]bool, len(s))
    max := 1
    ret := string(s[0])
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
        dp[i][i] = true
    }

    for j := 1; j < len(s); j++ {
        for i := 0; i < j; i++ {
            if s[i] != s[j] {
                dp[i][j] = false
            } else {
                if j-1-(i+1)+1 < 2 {
                    dp[i][j] = true
                } else {
                    dp[i][j] = dp[i+1][j-1]
                }
            }
            if dp[i][j] && j-i+1 > max {
                max = j - i + 1
                ret = s[i : j+1]
            }
        }
    }
    return ret
}
