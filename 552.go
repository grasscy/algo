package algo

var mod int = 1e9 + 7
var memo [][][]int

func checkRecord(n int) int {
	memo = make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		memo[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			memo[i][j] = make([]int, 3)
			for k := 0; k < 3; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	return d(n, 0, 0)
}

func d(n int, acnt int, endLLcnt int) int {

	if acnt >= 2 {
		return 0
	}
	if endLLcnt >= 3 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if memo[n][acnt][endLLcnt] != -1 {
		return memo[n][acnt][endLLcnt]
	}
	ans := d(n-1, acnt+1, 0) % mod
	ans += d(n-1, acnt, endLLcnt+1) % mod
	ans += d(n-1, acnt, 0) % mod
	memo[n][acnt][endLLcnt] = ans % mod
	return ans % mod
}
