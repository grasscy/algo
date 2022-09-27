package algo

func isValidSudoku(board [][]byte) bool {
	raw := make([]map[byte]bool, 10)
	col := make([]map[byte]bool, 10)
	ks := make([]map[byte]bool, 10)
	for i := 0; i < 10; i++ {
		raw[i] = map[byte]bool{}
		col[i] = map[byte]bool{}
		ks[i] = map[byte]bool{}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == '.' {
				continue
			}
			b := board[i][j] - '0'
			if raw[i][b] == true {
				return false
			}
			if col[j][b] == true {
				return false
			}
			k := 3*(i/3) + j/3
			if ks[k][b] == true {
				return false
			}

			raw[i][b] = true
			col[j][b] = true
			ks[k][b] = true
		}
	}
	return true
}
