package t1_199

var vis [][]bool

func exist(board [][]byte, word string) bool {
	vis = make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		vis[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if d(board, word, i, j, 0) {
				return true
			}
		}

	}
	return false
}

func d(board [][]byte, word string, ii int, ij int, wi int) bool {
	if wi == len(word) {
		return true
	}
	if ii < 0 || ii >= len(board) || ij < 0 || ij >= len(board[0]) || vis[ii][ij] {
		return false
	}
	if board[ii][ij] != word[wi] {
		return false
	}
	vis[ii][ij] = true
	ans := d(board, word, ii+1, ij, wi+1) || d(board, word, ii-1, ij, wi+1) || d(board, word, ii, ij+1, wi+1) || d(board, word, ii, ij-1, wi+1)
	if ans {
		return ans
	}
	vis[ii][ij] = false
	return false
}
