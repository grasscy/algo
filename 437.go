package algo

var count int

func pathSum(root *TreeNode, targetSum int) int {
	count = 0
	return sum(root, targetSum)
}

func sum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	d(root, targetSum)
	sum(root.Left, targetSum)
	sum(root.Right, targetSum)
	return count
}

func d(root *TreeNode, targetSum int) {
	if root == nil {
		return
	}
	if root.Val == targetSum {
		count++
	}
	d(root.Left, targetSum-root.Val)
	d(root.Right, targetSum-root.Val)
}
