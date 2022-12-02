package t1_199

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ans := ""

	n := 0
	//if len(num2) > len(num1) {
	//    num1, num2 = num2, num1
	//}
	for i := len(num2) - 1; i >= 0; i-- {
		i2 := m(num1, int(num2[i]-'0'))
		for j := 0; j < n; j++ {
			i2 += "0"
		}
		n++
		ans = add(ans, i2)
	}

	return ans
}

func add(n1 string, n2 string) string {
	jin := 0
	ans := ""

	ll := max(len(n1), len(n2))

	for i := 1; i <= ll; i++ {
		var v1, v2 int
		if len(n1)-i >= 0 {
			v1 = int(n1[len(n1)-i] - '0')
		} else {
			v1 = 0
		}
		if len(n2)-i >= 0 {
			v2 = int(n2[len(n2)-i] - '0')
		} else {
			v2 = 0
		}
		t := v1 + v2 + jin
		ans = strconv.Itoa(t%10) + ans
		jin = t / 10
	}

	if jin != 0 {
		ans = strconv.Itoa(jin) + ans
	}
	return ans
}

func m(num string, n int) string {
	ans := ""
	//bei := 1
	jin := 0

	for i := len(num) - 1; i >= 0; i-- {
		v := int(num[i]-'0') * n
		vv := v + jin
		ans = strconv.Itoa(vv%10) + ans
		jin = vv / 10
	}
	if jin != 0 {
		ans = strconv.Itoa(jin) + ans
	}
	return ans
}
