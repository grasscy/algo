package ac

var ans int

func diameterOfBinaryTree(root *TreeNode) int {
	ans = 0
	d(root)
	return ans
}

func d(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := d(root.Left)
	r := d(root.Right)

	ans = max(ans, l+r)

	return max(r, l) + 1
}
