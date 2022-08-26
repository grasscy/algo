package ac

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, v := range nums {
		m[v] = m[v] + 1
	}
	tmp := make([][]int, len(nums)+1)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = []int{}
	}

	for key, v := range m {
		tmp[v] = append(tmp[v], key)
	}

	r := []int{}
	sk := 0
	for i := len(nums); i >= 0; i-- {
		for _, v := range tmp[i] {
			if sk < k {
				r = append(r, v)
				sk++
			} else {
				return r
			}
		}

	}

	return r
}
