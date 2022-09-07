package ac

func checkInclusion(s1 string, s2 string) bool {
	needs := map[string]int{}
	for _, v := range s1 {
		needs[string(v)]++
	}
	left, right := 0, 0
	win := map[string]int{}
	valid := 0
	for right < len(s2) {
		c := string(s2[right])
		right++
		if _, ok := needs[c]; ok {
			win[c]++
			if win[c] == needs[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(needs) {
				return true
			}
			d := string(s2[left])
			left++
			if _, ok := needs[d]; ok {
				if win[d] == needs[d] {
					valid--
				}
				win[d]--
			}
		}
	}
	return false
}
