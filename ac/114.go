package ac

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left
	pre := left
	for pre.Right != nil {
		pre = pre.Right
	}
	pre.Right = right
}
