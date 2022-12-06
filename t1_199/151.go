package t1_199

import (
	"strings"
)

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	ans := []string{}
	for _, v := range split {
		if v != "" {
			ans = append(ans, v)
		}
	}
	if len(ans) == 1 {
		return ans[0]
	}
	ss := ans[len(ans)-1]
	for i := len(ans) - 2; i >= 0; i-- {
		ss = ss + " " + ans[i]
	}

	return ss
}
