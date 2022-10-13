package ac

import "math"

var memo [][]int

func minPathSum(grid [][]int) int {
	memo = initArray(len(grid), len(grid[0]))
	fill2(memo, -1)
	return dp(grid, len(grid)-1, len(grid[0])-1)
}

func dp(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt32
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}

	ans := min(dp(grid, i-1, j), dp(grid, i, j-1)) + grid[i][j]
	memo[i][j] = ans
	return ans
}
