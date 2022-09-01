package ac

func validateStackSequences(pushed []int, popped []int) bool {
	st := []int{}
	j := 0
	for _, v := range pushed {
		st = append(st, v)
		for len(st) > 0 && popped[j] == st[len(st)-1] {
			st = st[:len(st)-1]
			j++
		}
	}
	return len(st) == 0
}
