package t200_399

func numIslands(grid [][]byte) int {
	r := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				r++
				dfs(grid, i, j)
			}
		}
	}
	return r
}

func dfs(grid [][]byte, i int, j int) {
	grid[i][j] = '0'
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if i-1 >= 0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
	if j-1 >= 0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
}
