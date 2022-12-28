package ac

func minimumLength(s string) int {
	l, r := 0, len(s)-1
	for l < r && s[l] == s[r] {
		for ; l+1 < r && s[l] == s[l+1]; l++ {
		}
		for ; r-1 > l && s[r] == s[r-1]; r-- {
		}
		l++
		r--
	}
	return r - l + 1
}
