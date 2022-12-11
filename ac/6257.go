package ac

func deleteGreatestValue(grid [][]int) int {
	ans := 0
	n := len(grid[0])
	for n != 0 {
		mmax := 0
		n--
		for _, v := range grid {
			rmmax := 0
			for i, vv := range v {
				if vv > v[rmmax] {
					rmmax = i
				}
			}
			if mmax < v[rmmax] {
				mmax = v[rmmax]
			}
			v[rmmax] = -1
		}
		ans += mmax

	}

	return ans
}
