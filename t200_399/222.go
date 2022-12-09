package t200_399

var ans int

func countNodes(root *TreeNode) int {
	ans = 0
	d(root)
	return ans
}

func d(r *TreeNode) {
	if r == nil {
		return
	}
	ans++
	d(r.Left)
	d(r.Right)
}
