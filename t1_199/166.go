package t1_199

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	ans := ""

	if numerator*denominator < 0 {
		ans = "-"
	}
	numerator = abs(numerator)
	denominator = abs(denominator)

	t := numerator / denominator
	ans = strconv.Itoa(t)
	numerator = numerator % denominator

	if numerator == 0 {
		return ans
	}
	ans += "."
	yu := -1
	m := map[string]int{}
	for yu != 0 {
		numerator *= 10
		t := numerator / denominator
		yu = numerator % denominator
		kk := fmt.Sprintf("%d-%d", numerator, yu)
		if m[kk] != 0 {
			aa := ""
			for i := 0; i < m[kk]; i++ {
				aa += string(ans[i])
			}
			aa += "("
			for i := m[kk]; i < len(ans); i++ {
				aa += string(ans[i])
			}
			aa += ")"

			return aa
		}
		numerator = yu
		ans += strconv.Itoa(t)
		m[kk] = len(ans) - 1
	}
	return ans
}
