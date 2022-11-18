package t200_399

import "strconv"

func diffWaysToCompute(expression string) []int {
	atoi, err := strconv.Atoi(expression)
	if err == nil {
		return []int{atoi}
	}

	ans := []int{}

	for i, v := range expression {
		if v == '+' || v == '-' || v == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch v {
					case '+':
						ans = append(ans, l+r)
					case '-':
						ans = append(ans, l-r)
					case '*':
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	return ans

}
