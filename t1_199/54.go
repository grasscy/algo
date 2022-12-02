package t1_199

var dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
var vis [][]int

func spiralOrder(matrix [][]int) []int {
	ans := []int{}

	vis = initArray(len(matrix), len(matrix[0]))
	i, j := 0, 0
	k := 0
	for {
		ans = append(ans, matrix[i][j])
		vis[i][j] = 1

		if check(matrix, i, j, k) {
			i += dir[k][0]
			j += dir[k][1]
			continue
		}
		if check(matrix, i, j, (k+1)%4) {
			i += dir[(k+1)%4][0]
			j += dir[(k+1)%4][1]
			k++
			if k > 3 {
				k = 0
			}
			continue
		}
		return ans
	}
}

func check(matrix [][]int, i, j int, k int) bool {
	i += dir[k][0]
	j += dir[k][1]
	if i >= len(matrix) || i < 0 || j >= len(matrix[0]) || j < 0 || vis[i][j] == 1 {
		return false
	}
	return true
}
