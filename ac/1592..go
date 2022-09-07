package ac

func reorderSpaces(text string) string {
	n := 0
	s := ""
	is := []string{}
	for _, v := range text {
		if string(v) == " " {
			n++
			if s != "" {
				is = append(is, s)
				s = ""
			}
		} else {
			s += string(v)
		}
	}
	if s != "" {
		is = append(is, s)
	}
	ans := ""
	if len(is) == 1 {
		ans += is[0]
		for i := 0; i < n; i++ {
			ans += " "
		}
		return ans
	}
	nums := n / (len(is) - 1)
	last := n - nums*(len(is)-1)
	for j := 0; j < len(is)-1; j++ {
		ans += is[j]
		for i := 0; i < nums; i++ {
			ans += " "
		}
	}
	ans += is[len(is)-1]
	for i := 0; i < last; i++ {
		ans += " "
	}
	return ans

}
