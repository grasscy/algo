package jianzhioffer

var ans []string
var vis []bool
var m map[string]bool

func permutation(s string) []string {
	ans = []string{}
	vis = make([]bool, len(s))
	m = map[string]bool{}
	b(s, "")
	return ans
}

func b(s string, tmp string) {
	if len(s) == len(tmp) {
		if !m[tmp] {
			ans = append(ans, tmp)
		}
		m[tmp] = true
		return
	}
	for i := 0; i < len(s); i++ {
		if vis[i] {
			continue
		}
		vis[i] = true
		tmp += string(s[i])
		b(s, tmp)
		tmp = tmp[:len(tmp)-1]
		vis[i] = false
	}
}
