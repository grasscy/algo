package ac

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stk := []int{0}

	for i := 1; i < len(temperatures); i++ {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] < temperatures[i] {
			top := stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			ans[top] = i - top
		}
		stk = append(stk, i)
	}
	return ans
}
