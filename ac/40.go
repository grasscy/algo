package ac

import (
	"sort"
)

var ans [][]int
var vis []bool

func combinationSum2(candidates []int, target int) [][]int {
	ans = [][]int{}
	vis = make([]bool, len(candidates))
	sort.Ints(candidates)
	b(candidates, target, 0, []int{}, 0)
	return ans
}

func b(candidates []int, target int, sum int, cur []int, start int) {
	if sum == target {
		ans = append(ans, append([]int{}, cur...))
		return
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		if vis[i] || (i > start && candidates[i] == candidates[i-1]) {
			continue
		}
		vis[i] = true
		cur = append(cur, candidates[i])
		b(candidates, target, sum+candidates[i], cur, i+1)
		cur = cur[:len(cur)-1]
		vis[i] = false
	}
}
