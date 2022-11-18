package pass

func lengthOfLongestSubstring(s string) int {
	ans := 0
	count := map[uint8]int{}
	l, r := 0, 0
	for ; r < len(s); r++ {
		if count[s[r]] == 0 {
			ans = max(ans, r-l+1)
		} else {
			for ; l < r && count[s[r]] > 0; l++ {
				count[s[l]]--
			}
		}
		count[s[r]]++
	}
	return ans
}
