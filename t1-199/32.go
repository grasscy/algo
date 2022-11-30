package pass

func longestValidParentheses(s string) int {
	ans := 0
	stk := []int{-1}
	for i, v := range s {
		if v == '(' {
			stk = append(stk, i)
		} else {
			stk = stk[:len(stk)-1]
			if len(stk) == 0 {
				stk = append(stk, i)
			} else {
				ans = max(ans, i-stk[len(stk)-1])
			}
		}
	}
	return ans
}
