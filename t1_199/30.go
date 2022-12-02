package t1_199

func findSubstring(s string, words []string) []int {
	all := map[string]int{}
	for _, v := range words {
		all[v]++
	}

	ans := []int{}

	wl := len(words[0])
	n := len(words) * wl
	for i := 0; i < len(s)-n+1; i++ {
		f := true
		a := map[string]int{}
		for k, v := range all {
			a[k] = v
		}
		for j := i; j < i+n-wl+1; j += wl {
			ss := s[j : j+wl]
			a[ss]--
			if a[ss] < 0 {
				f = false
				break
			}
		}
		for _, v := range a {
			if v != 0 {
				f = false
				break
			}
		}
		if f {
			ans = append(ans, i)
		}

	}

	return ans
}
