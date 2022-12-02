package t1_199

var mm int
var ans string

func longestPalindrome(s string) string {
	mm = 1
	ans = string(s[0])
	for i := 0; i < len(s); i++ {
		f(s, i, i)
		f(s, i, i+1)
	}
	return ans
}

func f(s string, start int, end int) {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	a := s[start+1 : end]
	if len(a) > mm {
		mm = len(a)
		ans = a
	}
}
