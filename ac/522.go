package ac

func findLUSlength(strs []string) int {
	m := map[string]int{}
	for _, s := range strs {
		strings := finds(s)
		for _, str := range strings {
			m[str]++
		}
	}
	ans := -1
	for k, v := range m {
		if v == 1 {
			if len(k) > ans {
				ans = len(k)
			}
		}
	}
	return ans
}

func finds(s string) []string {
	m := map[string]bool{}
	m[""] = true
	for _, v := range s {
		t := []string{}
		for a, _ := range m {
			t = append(t, a)
		}
		for _, a := range t {
			m[a+string(v)] = true
			m[string(v)] = true
		}
	}
	ans := []string{}
	for a, _ := range m {
		ans = append(ans, a)
	}
	return ans
}
