package ac

func constructArr(a []int) []int {
	ans := make([]int, len(a))
	if len(a) == 0 {
		return ans
	}
	ans[0] = 1
	for i := 1; i < len(a); i++ {
		ans[i] = ans[i-1] * a[i-1]
	}
	r := 1
	for i := len(a) - 1; i >= 0; i-- {
		ans[i] = ans[i] * r
		r *= a[i]
	}
	return ans
}
