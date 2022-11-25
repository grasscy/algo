package pass

import "math"

func myAtoi(s string) int {
	f := 1
	ans := 0
	start := 0
	for ; start < len(s) && s[start] == ' '; start++ {

	}
	if start == len(s) {
		return 0
	}
	if s[start] != '-' && s[start] != '+' && (s[start] < '0' || s[start] > '9') {
		return 0
	}
	if s[start] == '-' {
		f = -1
		start++
	} else if s[start] == '+' {
		start++
	}
	for ; start < len(s); start++ {
		if s[start] >= '0' && s[start] <= '9' {
			v := s[start] - '0'
			ans = ans*10 + int(v)

			if ans*f > math.MaxInt32 {
				return math.MaxInt32
			}
			if ans*f < math.MinInt32 {
				return math.MinInt32
			}

		} else {
			break
		}
	}

	return ans * f
}
