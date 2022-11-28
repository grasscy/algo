package pass

import "math"

func divide(dividend int, divisor int) int {
	f := 1
	if dividend*divisor < 0 {
		f = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	ans := 0
	div, tmp := divisor, 1
	for dividend >= divisor {
		for dividend > div<<1 {
			div <<= 1
			tmp <<= 1
		}
		dividend -= div
		ans += tmp
		div = divisor
		tmp = 1
	}
	if f*ans > math.MaxInt32 {
		return math.MaxInt32
	}
	return f * ans
}
