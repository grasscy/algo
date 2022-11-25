package jianzhioffer

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	counts := map[string]int{}

	ans := 0
	for right < len(s) {
		c := string(s[right])
		right++

		counts[c]++
		if counts[c] == 1 && len(counts) > ans {
			ans = len(counts)
		} else {
			for counts[c] > 1 {
				d := string(s[left])
				left++
				counts[d]--
				if counts[d] == 0 {
					delete(counts, d)
				}
			}
		}
	}
	return ans
}
