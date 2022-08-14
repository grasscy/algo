package ac

func findDisappearedNumbers(nums []int) []int {
	r := []int{}
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			r = append(r, i)
		}
	}
	return r
}
