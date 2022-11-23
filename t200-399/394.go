package t200_399

import (
	"strconv"
)

func decodeString(s string) string {
	stkk := []int{}
	stks := []string{}
	ans := ""
	kk := ""
	for i := 0; i < len(s); i++ {
		v := s[i]
		vs := string(v)
		if '0' <= v && v <= '9' {
			kk += vs
		} else if 'a' <= v && v <= 'z' {
			ans += vs
		} else if v == '[' {
			atoi, _ := strconv.Atoi(kk)
			stkk = append(stkk, atoi)
			kk = ""
			stks = append(stks, ans)
			ans = ""
		} else {
			n := stkk[len(stkk)-1]
			stkk = stkk[:len(stkk)-1]
			t := ""
			for n != 0 {
				t += ans
				n--
			}
			ans = t
			if len(stks) > 0 {
				top := stks[len(stks)-1]
				stks = stks[:len(stks)-1]
				ans = top + ans
			}
		}
	}

	return ans
}
