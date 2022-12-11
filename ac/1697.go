package ac

import (
	"sort"
)

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if fa[x] == x {
			return x
		}
		fa[x] = find(fa[x])
		return fa[x]
	}

	merge := func(a, b int) {
		from, to := find(a), find(b)
		if from != to {
			fa[from] = to
		}
	}

	connected := func(a, b int) bool {
		return find(a) == find(b)
	}

	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool {
		return queries[id[i]][2] < queries[id[j]][2]
	})

	ans := make([]bool, len(queries))
	j := 0
	for _, i := range id {
		q := queries[i]
		for ; j < len(edgeList) && edgeList[j][2] < q[2]; j++ {
			merge(edgeList[j][0], edgeList[j][1])
		}
		ans[i] = connected(q[0], q[1])
	}
	return ans
}
