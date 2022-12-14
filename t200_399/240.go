package t200_399

func searchMatrix(matrix [][]int, target int) bool {
	i := 0
	j := len(matrix[0]) - 1

	for i < len(matrix) && j >= 0 {
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
