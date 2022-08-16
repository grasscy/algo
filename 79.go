package algo

var dir = [][]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
var rows int
var cols int
var word string
var board [][]byte
var visited [][]bool

func exist(board1 [][]byte, word1 string) bool {
    board = board1
    word = word1
    rows = len(board)
    cols = len(board[0])

    visited = make([][]bool, rows)
    for i := 0; i < rows; i++ {
        visited[i] = make([]bool, cols)
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if dfs(i, j, 0) {
                return true
            }
        }
    }
    return false
}

func dfs(x int, y int, begin int) bool {
    if begin == len(word)-1 {
        return board[x][y] == word[begin]
    }
    if board[x][y] == word[begin] {
        visited[x][y] = true
        for _, d := range dir {
            nx := x + d[0]
            ny := y + d[1]
            if in(nx, ny) {
                if !visited[nx][ny] && dfs(nx, ny, begin+1) {
                    return true
                }
            }
        }
        visited[x][y] = false
    }

    return false
}

func in(i, j int) bool {
    return i >= 0 && i < rows && j >= 0 && j < cols
}
