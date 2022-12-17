package t1_199

var g [][]bool

func minCut(s string) int {
    sl := len(s)
    dp := make([]int, sl)
    g = make([][]bool, sl)
    for i := range g {
        g[i] = make([]bool, sl)
    }

    for r := 0; r < sl; r++ {
        for l := r; l >= 0; l-- {
            if l == r {
                g[l][r] = true
            } else {
                if s[l] == s[r] {
                    if r-l == 1 || g[l+1][r-1] {
                        g[l][r] = true
                    }
                }
            }
        }
    }

    for i := 0; i < sl; i++ {
        if g[0][i] {
            dp[i] = 0
        } else {
            dp[i] = i
            for j := 0; j <= i; j++ {
                if g[j][i] {
                    if dp[i] > 1+dp[j-1] {
                        dp[i] = 1 + dp[j-1]
                    }
                }
            }
        }
    }
    return dp[sl-1]
}
