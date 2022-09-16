package ac

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 1
	l := 0
	m := map[uint8]bool{}
	count := 0
	for r := l; r < len(s); r++ {
		for m[s[r]] == true {
			delete(m, s[l])
			l++
			count--
		}
		m[s[r]] = true
		count++
		if count > ans {
			ans = count
		}
	}
	return ans
}
