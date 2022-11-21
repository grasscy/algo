package t200_399

import "strconv"

func decodeString(s string) string {
	res := ""
	stknum := []int{}
	stk := []string{}

	num := ""

	for _, v := range s {
		if '0' <= v && v <= '9' {
			num += string(v)
		} else if 'a' <= v && v <= 'z' {
			res += string(v)
		} else if v == '[' {
			atoi, _ := strconv.Atoi(num)
			stknum = append(stknum, atoi)
			stk = append(stk, res)
			num = ""
			res = ""
		} else if v == ']' {
			tmp := ""
			n := stknum[len(stknum)-1]
			stknum = stknum[:len(stknum)-1]
			for i := 0; i < n; i++ {
				tmp += res
			}
			last := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			res = last + tmp
		}

	}
	return res
}
