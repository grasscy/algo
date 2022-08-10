package algo

// opt
func maximalSquare(matrix [][]byte) int {
    dp := make([][]int, len(matrix))
    ret := 0
    for i := 0; i < len(matrix); i++ {
        dp[i] = make([]int, len(matrix[0]))
        if matrix[i][0] == '1' {
            dp[i][0] = 1
            ret = 1
        } else {
            dp[i][0] = 0
        }

    }
    for i := 0; i < len(matrix[0]); i++ {
        if matrix[0][i] == '1' {
            dp[0][i] = 1
            ret = 1
        } else {
            dp[0][i] = 0
        }
    }

    for i := 1; i < len(matrix); i++ {
        for j := 1; j < len(matrix[0]); j++ {
            if matrix[i][j] == '0' {
                dp[i][j] = 0
            } else {
                dp[i][j] = min(min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
                if dp[i][j] > ret {
                    ret = dp[i][j]
                }
            }
        }
    }
    return ret * ret
}
