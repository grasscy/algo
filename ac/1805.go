package ac

func numDifferentIntegers(word string) int {
	m := map[string]bool{}
	s := ""
	ans := 0
	for _, v := range word {
		if 'a' <= v && v <= 'z' {
			if len(s) > 0 {
				if !m[s] {
					ans++
				}
				m[s] = true
			}
			s = ""
		} else {
			if s == "0" {
				s = string(v)
			} else {
				s = s + string(v)
			}
		}
	}

	if len(s) > 0 {
		if !m[s] {
			ans++
		}
	}
	return ans
}
