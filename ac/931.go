package ac

import "math"

func minFallingPathSum(matrix [][]int) int {
    n := len(matrix)
    f := make([][]int, n)
    for i := 0; i < n; i++ {
        f[i] = make([]int, n)
        f[0][i] = matrix[0][i]
    }

    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            v := matrix[i][j]
            if j == 0 {
                f[i][j] = min(f[i-1][j], f[i-1][j+1]) + v
            } else if j == n-1 {
                f[i][j] = min(f[i-1][j], f[i-1][j-1]) + v
            } else {
                f[i][j] = min(min(f[i-1][j], f[i-1][j-1]), f[i-1][j+1]) + v
            }
        }
    }

    r := math.MaxInt32
    for i := 0; i < n; i++ {
        r = min(r, f[n-1][i])
    }
    return r
}
