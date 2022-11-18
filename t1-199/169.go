package pass

func majorityElement(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		n, ok := m[v]
		if !ok {
			m[v] = 1
		} else {
			m[v] = n + 1
		}
		n = m[v]
		if n > len(nums)/2 {
			return v
		}
	}
	return 0
}
