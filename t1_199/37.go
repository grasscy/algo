package t1_199

var ans [][]int
var f bool

func solveSudoku(board [][]byte) {
	ans = initArray(9, 9)
	f = false
	b(board, 0, 0)
}

func b(board [][]byte, m, n int) {
	if m == len(board)-1 && n == len(board[0]) {
		f = true
		return
	}
	if n == 9 {
		m++
		n = 0
	}
	if board[m][n] == '.' {
		for k := '1'; k <= '9'; k++ {
			if !check(board, m, n, byte(k)) {
				continue
			}
			board[m][n] = byte(k)
			b(board, m, n+1)
			if f {
				return
			}
			board[m][n] = '.'
		}
	} else {
		b(board, m, n+1)
	}
}

func check(board [][]byte, m, n int, k byte) bool {
	for i := 0; i < 9; i++ {
		if board[m][i] == k {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][n] == k {
			return false
		}
	}

	is := (m / 3) * 3
	js := (n / 3) * 3

	for i := is; i < is+3; i++ {
		for j := js; j < js+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
