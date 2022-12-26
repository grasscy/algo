package ac

func minimumBoxes(n int) (ans int) {
	maxN := 0
	for i := 1; maxN+ans+i <= n; i++ {
		ans += i
		maxN += ans
	}
	for j := 1; maxN < n; j++ {
		ans++
		maxN += j
	}
	return
}
