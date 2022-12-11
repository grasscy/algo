package ac

import (
	"sort"
)

var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxPoints(grid [][]int, queries []int) []int {
	m, n := len(grid), len(grid[0])
	mn := m * n

	fa := make([]int, mn)
	size := make([]int, mn)

	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	merge := func(from, to int) {
		from = find(from)
		to = find(to)
		if from != to {
			fa[from] = to
			size[to] += size[from]
		}
	}

	type tuple struct{ x, i, j int }
	a := make([]tuple, 0, mn)
	for i, row := range grid {
		for j, x := range row {
			a = append(a, tuple{x, i, j})
		}
	}
	sort.Slice(a, func(i, j int) bool { return a[i].x < a[j].x })

	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return queries[id[i]] < queries[id[j]] })

	ans := make([]int, len(queries))
	j := 0
	for _, i := range id {
		q := queries[i]
		for ; j < mn && a[j].x < q; j++ {
			x, y := a[j].i, a[j].j
			for _, d := range dir {
				x2, y2 := x+d[0], y+d[1]
				if 0 <= x2 && x2 < m && 0 <= y2 && y2 < n && grid[x2][y2] < q {
					merge(x*n+y, x2*n+y2)
				}
			}
		}
		if grid[0][0] < q {
			ans[i] = size[find(0)]
		}
	}
	return ans

}
