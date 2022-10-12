package ac

var count int
var ans int

func kthSmallest(root *TreeNode, k int) int {
	count = 0
	ans = -1

	tra(root, k)
	return ans
}

func tra(root *TreeNode, k int) {
	if root == nil {
		return
	}
	tra(root.Left, k)
	count++
	if count == k {
		ans = root.Val
	}
	tra(root.Right, k)
}
