package ac

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !isUnivalTree(root.Left) || !isUnivalTree(root.Right) {
		return false
	}
	if root.Left != nil && root.Left.Val != root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val != root.Val {
		return false
	}
	return true
}
