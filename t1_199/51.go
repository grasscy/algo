package t1_199

var ans [][]string
var vis [][]bool

func solveNQueens(n int) [][]string {
	ans = [][]string{}
	vis = make([][]bool, n)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, n)
	}
	b(n, 0, 0)
	return ans
}

func b(n int, cur int, m int) {
	if cur == n {
		tmp := []string{}
		for i := 0; i < n; i++ {
			s := ""
			for j := 0; j < n; j++ {
				if vis[i][j] {
					s += "Q"
				} else {
					s += "."
				}
			}
			tmp = append(tmp, s)
		}
		ans = append(ans, tmp)
		return
	}

	for i := m; i < n; i++ {
	a:
		for j := 0; j < n; j++ {
			dir := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
			for k := 0; k < n; k++ {
				if vis[i][k] {
					continue a
				}
				if vis[k][j] {
					continue a
				}
				for _, v := range dir {
					x := i + k*v[0]
					y := j + k*v[1]
					if x >= 0 && x < n && y >= 0 && y < n {
						if vis[x][y] {
							continue a
						}
					}
				}
			}
			vis[i][j] = true
			b(n, cur+1, i+1)
			vis[i][j] = false
		}
	}

}
