package t1_199

var ans [][]int

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans = [][]int{}
	if root == nil {
		return ans
	}
	b(root, targetSum, root.Val, []int{root.Val})
	return ans
}

func b(r *TreeNode, tg int, sum int, tmp []int) {
	if r == nil {
		return
	}
	if r.Left == nil && r.Right == nil {
		if tg == sum {
			ans = append(ans, append([]int{}, tmp...))
		}
		return
	}
	if r.Left != nil {
		b(r.Left, tg, sum+r.Left.Val, append(tmp, r.Left.Val))
	}
	if r.Right != nil {
		b(r.Right, tg, sum+r.Right.Val, append(tmp, r.Right.Val))
	}
}
