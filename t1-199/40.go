package pass

import (
	"sort"
)

var ans [][]int
var visit []bool

func combinationSum2(candidates []int, target int) [][]int {
	ans = [][]int{}
	visit = make([]bool, len(candidates))
	sort.Ints(candidates)
	b(candidates, target, []int{}, 0, 0)
	return ans
}

func b(candidates []int, target int, tmp []int, sum int, index int) {
	if sum == target {
		ans = append(ans, append([]int{}, tmp...))
		return
	}
	if sum > target {
		return
	}
	for i := index; i < len(candidates); i++ {
		if i > index && candidates[i] == candidates[i-1] {
			continue
		}
		sum += candidates[i]
		tmp = append(tmp, candidates[i])
		b(candidates, target, tmp, sum, i+1)
		tmp = tmp[:len(tmp)-1]
		sum -= candidates[i]
	}
}
