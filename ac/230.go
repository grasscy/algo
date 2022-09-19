package ac

var count int
var res int

func kthSmallest(root *TreeNode, k int) int {
	count = 0
	res = 0
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	count++
	if count == k {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}
