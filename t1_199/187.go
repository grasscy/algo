package t1_199

func findRepeatedDnaSequences(s string) []string {
	ans := []string{}
	if len(s) < 10 {
		return ans
	}
	exsit := map[string]int{}
	l, r := 0, 10
	for r <= len(s) {
		ss := s[l:r]
		exsit[ss]++
		l++
		r++
	}
	for k, v := range exsit {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
