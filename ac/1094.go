package ac

func carPooling(trips [][]int, capacity int) bool {
	diff := make([]int, 1111)
	diff[0] = 0
	for _, trip := range trips {
		diff[trip[1]] += trip[0]
		diff[trip[2]] -= trip[0]

	}
	t := 0
	for _, v := range diff {
		t += v
		if t > capacity {
			return false
		}
	}

	return true
}
