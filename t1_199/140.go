package t1_199

import "strings"

var ans []string

func wordBreak(s string, wordDict []string) []string {
	ans = []string{}
	wd := map[string]bool{}
	for _, v := range wordDict {
		wd[v] = true
	}
	b(s, wd, 0, []string{})
	return ans
}

func b(s string, wd map[string]bool, index int, tmp []string) {
	if index == len(s) {
		ans = append(ans, strings.Join(tmp, " "))
		return
	}
	for i := index; i < len(s); i++ {
		s2 := s[index : i+1]
		if wd[s2] {
			b(s, wd, i+1, append(tmp, s2))
		}
	}
}
