package t1_199

func isValid(s string) bool {
	ss := []int32{}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			ss = append(ss, v)
		} else if len(ss) == 0 {
			return false
		}

		if v == ')' {
			if ss[len(ss)-1] == '(' {
				ss = ss[:len(ss)-1]
			} else {
				ss = append(ss, v)
			}
		}
		if v == ']' {
			if ss[len(ss)-1] == '[' {
				ss = ss[:len(ss)-1]
			} else {
				ss = append(ss, v)
			}
		}
		if v == '}' {
			if ss[len(ss)-1] == '{' {
				ss = ss[:len(ss)-1]
			} else {
				ss = append(ss, v)
			}
		}
	}
	if len(ss) != 0 {
		return false
	}
	return true
}
