package pass

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	res := 0
	xx := x
	for xx/10 != 0 {
		res = res*10 + xx%10
		xx = xx / 10
	}
	res = res*10 + xx%10
	if res == x {
		return true
	}
	return false
}
