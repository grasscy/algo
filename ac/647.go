package ac

func countSubstrings(s string) int {
	ans := 0

	for i := 0; i < 2*len(s)-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < len(s) && s[l] == s[r] {
			ans++
			l--
			r++
		}
	}
	return ans
}
