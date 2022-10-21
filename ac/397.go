package ac

func integerReplacement(n int) int {
	ans := 0
	return d(n, ans)
}

func d(n int, ans int) int {
	if n == 1 {
		return ans
	}
	if n%2 == 0 {
		return d(n/2, ans+1)
	} else {
		return min(d(n-1, ans+1), d(n+1, ans+1))
	}
	return ans
}
