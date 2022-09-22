package ac

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	count := map[string]int{}
	ans := 0
	for ; r < len(s); r++ {
		c := string(s[r])
		for count[c] != 0 {
			lc := string(s[l])
			count[lc]--
			l++
		}
		count[c]++
		if r-l+1 > ans {
			ans = r - l + 1
		}

	}
	return ans
}
