package ac

import (
	"math"
	"strconv"
)

func printTree(root *TreeNode) [][]string {
	ans := [][]string{}
	m := findh(root) + 1
	n := int(math.Pow(2.0, float64(m))) - 1
	ans = make([][]string, m)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]string, n)
	}
	print(root, ans, 0, (n-1)/2, m-1)
	return ans
}

func print(root *TreeNode, ans [][]string, r int, c int, h int) {
	if root == nil {
		return
	}
	ans[r][c] = strconv.Itoa(root.Val)
	print(root.Left, ans, r+1, c-int(math.Pow(2.0, float64(h-r-1))), h)
	print(root.Right, ans, r+1, c+int(math.Pow(2.0, float64(h-r-1))), h)
}

func findh(root *TreeNode) int {
	if root == nil {
		return -1
	}
	return max(findh(root.Left), findh(root.Right)) + 1
}
