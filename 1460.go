package algo

func canBeEqual(target []int, arr []int) bool {
	m := map[int]int{}
	for _, v := range arr {
		_, ok := m[v]
		if ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	for _, v := range target {
		_, ok := m[v]
		if !ok {
			return false
		}
		m[v] -= 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
