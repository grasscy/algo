package ac

func invertTree(root *TreeNode) *TreeNode {
	f(root)
	return root
}

func f(r *TreeNode) {
	if r == nil {
		return
	}
	temp := r.Left
	r.Left = r.Right
	r.Right = temp
	f(r.Left)
	f(r.Right)
}
