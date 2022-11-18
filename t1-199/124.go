package pass

import "fmt"

var ans int

func maxPathSum(root *TreeNode) int {
	ans = root.Val
	d(root)
	return ans
}

func d(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := root.Val

	l := d(root.Left)
	r := d(root.Right)
	fmt.Println(root.Val, l, r)

	if l > 0 && r > 0 {
		ret += l + r
	} else if l < 0 && r < 0 {

	} else {
		ret += max(l, r)
	}
	ans = max(ans, ret)

	return max(root.Val, root.Val+max(l, r))
}
