package ac

import (
	"math"
)

func minScore(n int, roads [][]int) int {
	type edge struct{ to, d int }
	g := make([][]edge, n+1)
	for _, e := range roads {
		x, y, d := e[0], e[1], e[2]
		g[x] = append(g[x], edge{y, d})
		g[y] = append(g[y], edge{x, d})
	}
	ans := math.MaxInt32
	vis := make([]bool, n+1)
	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true
		for _, e := range g[x] {
			ans = min(ans, e.d)
			if !vis[e.to] {
				dfs(e.to)
			}
		}
	}
	dfs(1)
	return ans
}
