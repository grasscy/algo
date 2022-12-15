package ac

import "strconv"

func getLucky(s string, k int) int {
	nn := ""
	for _, v := range s {
		nn += strconv.Itoa(int(v - 'a' + 1))
	}

	t := 0
	for k != 0 {
		k--
		t = 0
		for _, v := range nn {
			t += int(v - '0')
		}
		nn = strconv.Itoa(t)
	}
	return t
}
