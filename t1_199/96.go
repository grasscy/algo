package t1_199

var memo [][]int

func numTrees(n int) int {
	memo = initArray(n+1, n+1)
	return c(1, n)
}

func c(lo, hi int) int {
	if lo >= hi {
		return 1
	}
	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}
	ans := 0
	for i := lo; i <= hi; i++ {
		ans += c(lo, i-1) * c(i+1, hi)
	}
	memo[lo][hi] = ans
	return ans
}
