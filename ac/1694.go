package ac

import "strings"

func reformatNumber(number string) string {
	ans := []string{}
	tmp := ""
	for i := 0; i < len(number); i++ {
		if '0' <= number[i] && number[i] <= '9' {
			tmp += string(number[i])
			if len(tmp) == 3 {
				ans = append(ans, tmp)
				tmp = ""
			}
		}
	}
	if len(tmp) == 1 {
		last := ans[len(ans)-1]
		ans = ans[:len(ans)-1]
		ans = append(ans, last[:2], string(last[2])+tmp)
	} else if len(tmp) != 0 {
		ans = append(ans, tmp)
	}
	return strings.Join(ans, "-")
}
