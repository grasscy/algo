package ac

import (
    "math"
    "strconv"
)

var res [][]string
var h int

func printTree(root *TreeNode) [][]string {
    h = geth(root) - 1
    m, n := h+1, int(math.Pow(2.0, float64(h+1.0)))-1
    res = make([][]string, m)
    for i := 0; i < m; i++ {
        res[i] = make([]string, n)
    }
    res[0][(n-1)/2] = strconv.Itoa(root.Val)
    set(0, (n-1)/2, root)
    return res
}

func set(r int, c int, node *TreeNode) {
    if node == nil {
        return
    }
    if node.Left != nil {
        res[r+1][c-int(math.Pow(2.0, float64(h-r-1)))] = strconv.Itoa(node.Left.Val)
        set(r+1, c-int(math.Pow(2.0, float64(h-r-1))), node.Left)
    }
    if node.Right != nil {
        res[r+1][c+int(math.Pow(2.0, float64(h-r-1)))] = strconv.Itoa(node.Right.Val)
        set(r+1, c+int(math.Pow(2.0, float64(h-r-1))), node.Right)
    }
}

func geth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    h = max(geth(root.Left), geth(root.Right)) + 1
    return h
}
