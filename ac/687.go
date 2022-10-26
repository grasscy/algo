package ac

var ans int

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans = 0
	longest(root)
	return ans
}

func longest(root *TreeNode) int {
	if root == nil {
		return 0
	}
	d(root)
	longest(root.Left)
	longest(root.Right)
	return ans
}

func d(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := 0, 0
	if root.Left != nil && root.Val == root.Left.Val {
		l = d(root.Left) + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		r = d(root.Right) + 1
	}
	ans = max(ans, l+r)
	return max(l, r)
}
