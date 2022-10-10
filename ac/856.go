package ac

import (
	"strconv"
)

func scoreOfParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	stk := []string{}
	for _, v := range s {
		c := string(v)
		switch c {
		case "(":
			stk = append(stk, c)
		case ")":
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			if top == "(" {
				stk = append(stk, "1")
			} else {
				sum := 0
				for top != "(" {
					n, _ := strconv.Atoi(top)
					sum += n
					top = stk[len(stk)-1]
					stk = stk[:len(stk)-1]
				}
				sum *= 2
				stk = append(stk, strconv.Itoa(sum))
			}
		}
	}
	var ans int
	for _, v := range stk {
		ss, _ := strconv.Atoi(v)
		ans += ss
	}
	return ans
}
