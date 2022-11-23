package ac

var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	d(root)
	return root
}

func d(root *TreeNode) {
	if root == nil {
		return
	}
	d(root.Right)
	sum += root.Val
	root.Val = sum
	d(root.Left)
}
