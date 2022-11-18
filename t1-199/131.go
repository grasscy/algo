package pass

var ans [][]string

func partition(s string) [][]string {
	ans = [][]string{}
	b(s, []string{}, 0)
	return ans
}

func b(s string, tmp []string, index int) {
	if index == len(s) {
		ans = append(ans, append([]string{}, tmp...))
		return
	}
	for i := index; i < len(s); i++ {
		if check(s, index, i) {
			tmp = append(tmp, s[index:i+1])
			b(s, tmp, i+1)
			tmp = tmp[:len(tmp)-1]
		}
	}

}

func check(s string, l, r int) bool {
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
