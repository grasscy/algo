package t1_199

var memo map[int]int

func climbStairs(n int) int {
	memo = map[int]int{}
	return f(n)
}

func f(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = f(n-1) + f(n-2)
	return memo[n]
}
