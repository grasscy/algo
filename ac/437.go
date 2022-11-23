package ac

var count int

func pathSum(root *TreeNode, targetSum int) int {
	count = 0
	p(root, targetSum)
	return count
}

func p(root *TreeNode, tg int) {
	if root == nil {
		return
	}
	d(root, tg)
	p(root.Left, tg)
	p(root.Right, tg)
}

func d(root *TreeNode, tg int) {
	if root == nil {
		return
	}
	if root.Val == tg {
		count++
	}
	d(root.Left, tg-root.Val)
	d(root.Right, tg-root.Val)
}
