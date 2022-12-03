package t1_199

func isPalindrome(s string) bool {
	ss := ""
	for _, v := range s {
		if v >= 'A' && v <= 'Z' {
			ss += string(v - 'A' + 'a')
		} else if v >= 'a' && v <= 'z' {
			ss += string(v)
		} else if v >= '0' && v <= '9' {
			ss += string(v)
		}
	}
	for i := 0; i < len(ss)/2; i++ {
		if ss[i] != ss[len(ss)-1-i] {
			return false
		}
	}
	return true

}
