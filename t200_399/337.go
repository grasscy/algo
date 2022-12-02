package t200_399

var memo map[*TreeNode]int

func rob(root *TreeNode) int {
	memo = map[*TreeNode]int{}
	return r(root)
}

func r(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if v, ok := memo[root]; ok {
		return v
	}
	mm := root.Val
	if root.Left != nil {
		mm += r(root.Left.Left) + r(root.Left.Right)
	}
	if root.Right != nil {
		mm += r(root.Right.Left) + r(root.Right.Right)
	}
	mm2 := r(root.Left) + r(root.Right)
	ans := max(mm, mm2)
	memo[root] = ans
	return ans
}
