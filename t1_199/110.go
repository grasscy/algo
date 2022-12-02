package t1_199

import "math"

func isBalanced(root *TreeNode) bool {
	return d(root)
}

func d(r *TreeNode) bool {
	if r == nil {
		return true
	}
	i := h(r.Left)
	j := h(r.Right)
	if math.Abs(float64(i-j)) > 1 {
		return false
	}
	return d(r.Left) && d(r.Right)
}

func h(r *TreeNode) int {
	if r == nil {
		return 0
	}
	return max(h(r.Left), h(r.Right)) + 1
}
