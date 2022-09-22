package ac

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i, v := range nums {
		if vv, ok := m[v]; ok {
			if i-vv <= k {
				return true
			}
		}
		m[v] = i

	}
	return false
}
