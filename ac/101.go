package ac

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return c(root.Left, root.Right)
}

func c(l, r *TreeNode) bool {
	if l == nil && r != nil {
		return false
	}
	if l != nil && r == nil {
		return false
	}
	if l == nil && r == nil {
		return true
	}
	if l.Val != r.Val {
		return false
	}
	return c(l.Left, r.Right) && c(l.Right, r.Left)
}
