package t1_199

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	right := root.Right
	root.Right = root.Left
	root.Left = nil

	r := root
	for r.Right != nil {
		r = r.Right
	}
	r.Right = right
	return
}
