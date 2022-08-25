package ac

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix[0]) - 1

	for j >= 0 && i < len(matrix) {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}

	}

	return false
}
