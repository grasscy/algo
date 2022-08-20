package ac

var r int

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth(root)
	return r
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := depth(root.Right)
	right := depth(root.Left)
	r = max(r, l+right+1)
	return max(l, right) + 1
}
