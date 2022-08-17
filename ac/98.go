package ac

import "math"

func isValidBST(root *TreeNode) bool {
    return f(root, math.MinInt64, math.MaxInt64)
}

func f(r *TreeNode, low, hi int) bool {
    if r == nil {
        return true
    }
    if r.Val <= low || r.Val >= hi {
        return false
    }
    return f(r.Left, low, r.Val) && f(r.Right, r.Val, hi)
}
