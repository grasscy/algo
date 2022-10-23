package algo

var m map[int]bool

func findTarget(root *TreeNode, k int) bool {
	m = map[int]bool{}
	return d(root, k)
}

func d(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	if m[k-root.Val] {
		return true
	}
	m[root.Val] = true
	return d(root.Left, k) || d(root.Right, k)
}
