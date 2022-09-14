package ac

import "strconv"

var m map[string]int
var ans []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ans = []*TreeNode{}
	m = map[string]int{}
	traverse(root)
	return ans
}

func traverse(node *TreeNode) string {
	res := ""
	if node == nil {
		return ""
	}
	res += strconv.Itoa(node.Val)
	res += "," + traverse(node.Left)
	res += "," + traverse(node.Right)

	if v, _ := m[res]; v == 1 {
		ans = append(ans, node)
	}
	m[res]++
	return res
}
