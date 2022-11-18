package pass

func setZeroes(matrix [][]int) {
	si := map[int]bool{}
	sj := map[int]bool{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				si[i] = true
				sj[j] = true
			}
		}
	}
	for k, _ := range si {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[k][j] = 0
		}
	}
	for k, _ := range sj {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[j][k] = 0
		}
	}
}
