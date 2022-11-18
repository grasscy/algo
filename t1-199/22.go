package pass

var r []string

func generateParenthesis(n int) []string {
	r = []string{}
	if n == 0 {
		return r
	}
	ff(n, "", 0, 0)
	return r
}

//opt
func ff(n int, s string, open int, close int) {
	if len(s) == n*2 {
		r = append(r, s)
		return
	}
	if open < n {
		s += "("
		ff(n, s, open+1, close)
		s = s[:len(s)-1]
	}
	if close < open {
		s += ")"
		ff(n, s, open, close+1)
		s = s[:len(s)-1]
	}

}
