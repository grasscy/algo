package t1_199

var ans int

func sumNumbers(root *TreeNode) int {
	ans = 0
	d(root, 0)
	return ans
}

func d(r *TreeNode, tmp int) {
	if r == nil {
		return
	}
	tmp = 10*tmp + r.Val
	if r.Left == nil && r.Right == nil {
		ans += tmp
		return
	}
	d(r.Left, tmp)
	d(r.Right, tmp)

}
