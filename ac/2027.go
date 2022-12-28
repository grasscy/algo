package ac

func minimumMoves(s string) int {
	ans := 0
	for i := 0; i < len(s); {
		if s[i] == 'X' {
			ans++
			i += 3
		} else {
			i++
		}
	}
	return ans
}
