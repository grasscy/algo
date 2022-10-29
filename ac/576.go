package ac

var memo [][][]int

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	memo = make([][][]int, m)
	for i := 0; i < len(memo); i++ {
		memo[i] = initArray(n, maxMove+1)
		fill2(memo[i], -1)
	}
	return find(m, n, maxMove, startRow, startColumn)
}

func find(m int, n int, maxMove int, startRow int, startColumn int) int {

	if startRow < 0 || startRow >= m || startColumn < 0 || startColumn >= n {
		return 1
	}
	if maxMove == 0 {
		return 0
	}
	if memo[startRow][startColumn][maxMove] != -1 {
		return memo[startRow][startColumn][maxMove]
	}
	ans := 0
	ans += find(m, n, maxMove-1, startRow-1, startColumn)
	ans += find(m, n, maxMove-1, startRow+1, startColumn)
	ans += find(m, n, maxMove-1, startRow, startColumn-1)
	ans += find(m, n, maxMove-1, startRow, startColumn+1)
	memo[startRow][startColumn][maxMove] = ans % (10e9 + 7)
	return ans % (10e9 + 7)
}
