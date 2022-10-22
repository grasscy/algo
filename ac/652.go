package ac

import "strconv"

var m map[string]int
var ans []*TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	m = map[string]int{}
	ans = []*TreeNode{}
	d(root)
	return ans
}

func d(r *TreeNode) string {
	if r == nil {
		return ","
	}
	res := strconv.Itoa(r.Val) + ","
	res += d(r.Left)
	res += d(r.Right)
	if m[res] == 1 {
		ans = append(ans, r)
	}
	m[res]++
	return res
}
