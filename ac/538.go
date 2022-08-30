package ac

var sum int

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	dfs(root)
	return root
}

func dfs(node *TreeNode) {
	if node == nil {
		return
	}
	dfs(node.Right)
	sum += node.Val
	node.Val = sum
	dfs(node.Left)
}
