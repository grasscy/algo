package ac

func closedIsland(grid [][]int) int {
	for j := 0; j < len(grid[0]); j++ {
		dfs(grid, 0, j)
		dfs(grid, len(grid)-1, j)
	}
	for i := 0; i < len(grid); i++ {
		dfs(grid, i, 0)
		dfs(grid, i, len(grid[0])-1)
	}

	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] == 1 {
		return
	}
	grid[i][j] = 1
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}
