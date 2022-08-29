package ac

func findAnagrams(s string, p string) []int {
	need := map[string]int{}
	for _, v := range p {
		need[string(v)]++
	}
	win := map[string]int{}

	left, right := 0, 0
	ans := []int{}

	valid := 0
	for ; right < len(s); right++ {
		c := string(s[right])
		if _, ok := need[c]; ok {
			win[c]++
			if win[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left+1 == len(p) {
				ans = append(ans, left)
			}

			d := string(s[left])
			left++
			if _, ok := need[d]; ok {
				if win[d] == need[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return ans
}
