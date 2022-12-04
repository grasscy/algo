package ac

import (
	"math"
	"sort"
)

var ans int
var nn map[int]int

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	ans = math.MaxInt32
	sort.Ints(toppingCosts)
	for _, v := range baseCosts {
		nn = map[int]int{}
		b(toppingCosts, target, v, 0)
	}
	return ans
}

func b(toppingCosts []int, tg int, sum int, index int) {
	if abs(tg-ans) > abs(tg-sum) {
		ans = sum
	} else if abs(tg-ans) == abs(tg-sum) && ans > sum {
		ans = sum
	}
	if sum > tg {
		return
	}

	for i := index; i < len(toppingCosts); i++ {
		b(toppingCosts, tg, sum+toppingCosts[i], i+1)
		b(toppingCosts, tg, sum+2*toppingCosts[i], i+1)
	}

}
