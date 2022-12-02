package t1_199

import "math"

func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(root *TreeNode, lo int, hi int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lo || root.Val >= hi {
		return false
	}
	return check(root.Left, lo, root.Val) && check(root.Right, root.Val, hi)

}
