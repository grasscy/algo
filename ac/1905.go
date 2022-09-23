package ac

var check bool
var g [][]int

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	g = grid1
	ans := 0
	for i := 0; i < len(grid2); i++ {
		for j := 0; j < len(grid2[0]); j++ {
			if grid2[i][j] == 1 {
				check = true
				dfs(grid2, i, j)
				if check {
					ans++
				}
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
	if grid[i][j] == 0 {
		return
	}
	if g[i][j] == 0 {
		check = false
	}
	grid[i][j] = 0
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
}
