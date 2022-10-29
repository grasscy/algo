package ac

import (
	"math"
)

var memo [][]int

func getMoneyAmount(n int) int {
	memo = initArray(n+1, n+1)
	return d(1, n)
}

func d(low int, high int) int {
	if low >= high {
		return 0
	}

	if memo[low][high] != 0 {
		return memo[low][high]
	}

	ret := math.MaxInt32
	for i := low; i <= high; i++ {
		t := max(d(low, i-1), d(i+1, high)) + i
		ret = min(ret, t)
	}
	memo[low][high] = ret
	return ret
}
