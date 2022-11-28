package pass

func myPow(x float64, n int) float64 {
	f := true
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		f = false
		n = -n
	}
	ans := 1.0
	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}

	if !f {
		ans = 1 / ans
	}
	return ans
}
