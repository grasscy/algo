package pass

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n-1
	for lo <= hi {
		mid := (lo + hi) / 2
		t := matrix[mid/n][mid%n]
		if t == target {
			return true
		} else if t > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}

	}
	return false

}
