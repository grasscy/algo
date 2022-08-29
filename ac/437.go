package ac

var sm map[int]int

func pathSum(root *TreeNode, targetSum int) int {
	sm = map[int]int{}
	sm[0] = 1
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, cursum int, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	cursum += root.Val

	res += sm[cursum-targetSum]

	sm[cursum]++
	res += dfs(root.Left, cursum, targetSum)
	res += dfs(root.Right, cursum, targetSum)
	sm[cursum]--

	return res
}
