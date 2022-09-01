package ac

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	st := []int{}
	for i := 0; i < len(temperatures); i++ {
		for len(st) > 0 && temperatures[i] > temperatures[st[len(st)-1]] {
			ans[st[len(st)-1]] = i - st[len(st)-1]
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return ans
}
