package ac

var ans int

func longestUnivaluePath(root *TreeNode) int {
	ans = 0

	dfs(root)

	return ans
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := dfs(node.Left)
	r := dfs(node.Right)

	l1, r1 := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		l1 = l + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		r1 = r + 1
	}

	ans = max(ans, l1+r1)
	return max(l1, r1)
}
