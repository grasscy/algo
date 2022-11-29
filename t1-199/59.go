package pass

func generateMatrix(n int) [][]int {
	ans := initArray(n, n)

	nn := 1
	u, d, l, r := 0, n-1, 0, n-1
	for true {
		for i := l; i <= r; i++ {
			ans[u][i] = nn
			nn++
		}
		u++
		if u > d {
			break
		}

		for i := u; i <= d; i++ {
			ans[i][r] = nn
			nn++
		}
		r--
		if l > r {
			break
		}

		for i := r; i >= l; i-- {
			ans[d][i] = nn
			nn++
		}
		d--
		if u > d {
			break
		}

		for i := d; i >= u; i-- {
			ans[i][l] = nn
			nn++
		}
		l++
		if l > r {
			break
		}

	}
	return ans
}
