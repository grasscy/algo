package t1_199

var ans []int

func postorderTraversal(root *TreeNode) []int {
	ans = []int{}
	b(root)
	return ans
}

func b(root *TreeNode) {
	if root == nil {
		return
	}
	b(root.Left)
	b(root.Right)
	ans = append(ans, root.Val)
}
