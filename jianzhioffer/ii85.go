package jianzhioffer

var ans []string

func generateParenthesis(n int) []string {
	ans = []string{}
	b(n, 0, 0, "")
	return ans
}

func b(n int, l, r int, str string) {
	if len(str) == 2*n {
		ans = append(ans, str)
	}
	if l < n {
		str += "("
		b(n, l+1, r, str)
		str = str[:len(str)-1]
	}
	if r < l {
		str += ")"
		b(n, l, r+1, str)
		str = str[:len(str)-1]
	}

}
