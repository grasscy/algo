package ac

var ans int
var visited [][]bool
var dir = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func getMaximumGold(grid [][]int) int {
	ans = 0
	visited = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			dfs(grid, i, j, 0)
		}
	}

	return ans
}

func dfs(grid [][]int, row, col int, tmp int) {
	if tmp > ans {
		ans = tmp
	}

	if !conti(grid, row, col) {
		return
	}

	visited[row][col] = true
	for _, d := range dir {
		i := row + d[0]
		j := col + d[1]

		dfs(grid, i, j, tmp+grid[row][col])
	}
	visited[row][col] = false
}

func conti(grid [][]int, row, col int) bool {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
		return false
	}

	if visited[row][col] {
		return false
	}

	if grid[row][col] == 0 {
		return false
	}
	return true
}
