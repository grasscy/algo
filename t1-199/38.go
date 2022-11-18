package pass

import "strconv"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)
	//s := "21"

	ans := ""
	c := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] != c {
			ans += strconv.Itoa(count)
			ans += string(c)
			c = s[i]
			count = 1
		} else {
			count++
		}
	}
	ans += strconv.Itoa(count)
	ans += string(c)
	return ans

}
