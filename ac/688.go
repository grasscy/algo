package ac

var memo [][][]float64

func knightProbability(n int, k int, row int, column int) float64 {
	memo = make([][][]float64, k+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]float64, n)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = make([]float64, n)
			for kk := 0; kk < len(memo[i][j]); kk++ {
				memo[i][j][kk] = -1
			}
		}
	}
	return dd(n, k, row, column)
}

func dd(n int, k int, row int, column int) float64 {
	if k == 0 {
		return 1.0
	}
	if memo[k][row][column] != -1 {
		return memo[k][row][column]
	}
	p := 0.0
	ds := [][]int{{-1, -2}, {1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}}
	for _, d := range ds {
		nx := row + d[0]
		ny := column + d[1]
		if nx >= 0 && nx < n && ny >= 0 && ny < n {
			p += dd(n, k-1, nx, ny) / 8.0
		}
	}
	memo[k][row][column] = p
	return p
}
