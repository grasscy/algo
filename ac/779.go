package ac

func kthGrammar(n int, k int) int {
	if n == 1 {
		return 0
	}
	above := kthGrammar(n-1, (k+1)/2)
	if k%2 == 0 {
		if above == 0 {
			return 1
		} else {
			return 0
		}
	}
	return above
}
