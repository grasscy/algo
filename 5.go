package algo

func longestPalindrome(s string) string {
	ans := string(s[0])
	for i, _ := range s {
		i2, i3 := ex(s, i, i)
		i4, i5 := ex(s, i, i+1)
		if i3-i2 > len(ans) {
			ans = s[i2+1 : i3]
		}
		if i5-i4 > len(ans) {
			ans = s[i4+1 : i5]
		}
	}

	return ans
}

func ex(s string, i1, i2 int) (int, int) {
	for i1 >= 0 && i2 < len(s) {
		if s[i1] == s[i2] {
			i1--
			i2++
		} else {
			break
		}
	}
	return i1, i2
}
