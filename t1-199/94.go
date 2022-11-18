package pass

var ans []int

func inorderTraversal(root *TreeNode) []int {
	ans = []int{}
	return d(root)
}

func d(root *TreeNode) {
	if root == nil {
		return
	}
	d(root.Left)
	ans = append(ans, root.Val)
	d(root.Right)
}
