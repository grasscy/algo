package algo

import "fmt"

var m map[int]int
var memo map[string]bool

func canCross(stones []int) bool {
	m = map[int]int{}
	for i, v := range stones {
		m[v] = i
	}
	memo = map[string]bool{}
	if stones[1] != 1 {
		return false
	}
	return d(stones, 1, 1)
}

func d(stones []int, c int, k int) bool {
	key := fmt.Sprintf("%d_%d", c, k)
	if _, ok := memo[key]; ok {
		return memo[key]
	}

	if len(stones)-1 <= c {
		return true
	}
	for i := -1; i <= 1; i++ {
		if k+i == 0 {
			continue
		}
		tg := stones[c] + k + i
		if m[tg] != 0 {
			if d(stones, m[tg], k+i) {
				memo[key] = true
				return true
			}
		}
	}
	memo[key] = false
	return false
}
