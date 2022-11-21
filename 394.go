package algo

func decodeString(s string) string {
	stkk := []int{}
	stks := []string{}
	ans := ""
	kk := ""
	for i := 0; i < len(s); i++ {
		v := s[i]
		vs := string(v)
		if '0' <= v && v <= '9' {
			kk += vs
		} else if 'a' <= v && v <= 'z' {
			ans += vs
		} else if v == '[' {

		} else {

		}
	}

	return ans
}
