package ac

var visited [][]bool

func exist(board [][]byte, word string) bool {
	visited = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if b(board, word, "", i, j, 0) {
				return true
			}
		}
	}
	return false
}

func b(board [][]byte, word string, tmp string, row, col int, index int) bool {
	if len(tmp) >= len(word) {
		return true
	}

	if row < 0 || row >= len(board) || col < 0 || col >= len(board[0]) {
		return false
	}
	if visited[row][col] {
		return false
	}
	if board[row][col] != word[index] {
		return false
	}
	visited[row][col] = true
	tmp += string(board[row][col])
	if b(board, word, tmp, row-1, col, index+1) {
		return true
	}
	if b(board, word, tmp, row+1, col, index+1) {
		return true
	}
	if b(board, word, tmp, row, col+1, index+1) {
		return true
	}
	if b(board, word, tmp, row, col-1, index+1) {
		return true
	}
	tmp = tmp[:len(tmp)-1]
	visited[row][col] = false
	return false
}
