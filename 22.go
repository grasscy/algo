package algo

var ans []string

func generateParenthesis(n int) []string {
	ans = []string{}
	d(n, "", 0, 0)
	return ans
}

func d(n int, tmp string, l, r int) {
	if l+r == 2*n {
		ans = append(ans, tmp)
		return
	}
	if l < n {
		d(n, tmp+"(", l+1, r)
	}
	if r < l {
		d(n, tmp+")", l, r+1)
	}
}
