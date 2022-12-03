package t1_199

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	i, j := 0, 0
	for ; j < n; j++ {
		d(board, i, j)
	}
	j--
	for ; i < m; i++ {
		d(board, i, j)
	}
	i--
	for ; j >= 0; j-- {
		d(board, i, j)
	}
	j++
	for ; i > 0; i-- {
		d(board, i, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 0 {
				board[i][j] = 'O'
			}
		}
	}

}

func d(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = 0
		d(board, i, j+1)
		d(board, i, j-1)
		d(board, i+1, j)
		d(board, i-1, j)
	}
}
