package pass

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ss := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		ss[i] = []string{}
	}
	r := 0
	f := true
	for i := 0; i < len(s); i++ {
		if r == numRows-1 {
			f = false
		}
		if r == 0 {
			f = true
		}
		ss[r] = append(ss[r], string(s[i]))
		if f {
			r++
		} else {
			r--
		}
	}
	ans := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(ss[i]); j++ {
			ans += ss[i][j]
		}
	}
	return ans

}
