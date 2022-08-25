package ac

func rob(root *TreeNode) int {
	ints := rr(root)
	return max(ints[0], ints[1])
}

func rr(root *TreeNode) []int {
	r := make([]int, 2)
	if root == nil {
		return r
	}
	left := rr(root.Left)
	right := rr(root.Right)

	r[0] = max(left[0], left[1]) + max(right[0], right[1])
	r[1] = left[0] + right[0] + root.Val
	return r
}

//
//var memo map[*TreeNode]int
//
//func rob(root *TreeNode) int {
//    memo = map[*TreeNode]int{}
//    return rr(root)
//}
//
//func rr(root *TreeNode) int {
//    if root == nil {
//        return 0
//    }
//    if v, ok := memo[root]; ok {
//        return v
//    }
//
//    money := root.Val
//    if root.Left != nil {
//        money += rob(root.Left.Left) + rob(root.Left.Right)
//    }
//    if root.Right != nil {
//        money += rob(root.Right.Left) + rob(root.Right.Right)
//    }
//
//    res := max(rob(root.Left)+rob(root.Right), money)
//    memo[root] = res
//    return res
//}
