package ac

import "strconv"

func tree2str(root *TreeNode) string {
	return d(root)
}

func d(root *TreeNode) string {
	if root == nil {
		return ""
	}
	tmp := strconv.Itoa(root.Val)
	if root.Left != nil {
		tmp += "("
		tmp += d(root.Left)
		tmp += ")"
	}
	if root.Right != nil {
		if root.Left == nil {
			tmp += "()"
		}
		tmp += "("
		tmp += d(root.Right)
		tmp += ")"
	}

	return tmp
}
