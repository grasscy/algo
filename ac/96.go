package ac

var memo [][]int

func numTrees(n int) int {
	memo = make([][]int, n+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([]int, n+1)
	}

	return count(1, n)
}

func count(lo int, hi int) int {
	if lo >= hi {
		return 1
	}

	if memo[lo][hi] != 0 {
		return memo[lo][hi]
	}

	ans := 0
	for i := lo; i <= hi; i++ {
		left := count(lo, i-1)
		right := count(i+1, hi)
		ans += left * right
	}
	memo[lo][hi] = ans
	return ans
}
