package t1_199

import "strconv"

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	ans := ""
	jin := 0
	for i := 0; i < len(a); i++ {
		aa := int(a[len(a)-1-i] - '0')
		bb := 0
		if len(b)-1-i >= 0 {
			bb = int(b[len(b)-1-i] - '0')
		}
		v := aa + bb + jin
		ans = strconv.Itoa(v%2) + ans
		jin = v / 2
	}
	if jin != 0 {
		ans = "1" + ans
	}
	return ans
}
