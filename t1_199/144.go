package t1_199

var ans []int

func preorderTraversal(root *TreeNode) []int {
	ans = []int{}
	b(root)
	return ans
}

func b(root *TreeNode) {
	if root == nil {
		return
	}
	ans = append(ans, root.Val)
	b(root.Left)
	b(root.Right)
}
