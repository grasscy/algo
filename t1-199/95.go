package pass

func generateTrees(n int) []*TreeNode {
	return generate(1, n)
}

func generate(lo int, hi int) []*TreeNode {
	ans := []*TreeNode{}
	if lo > hi {
		ans = append(ans, nil)
		return ans
	}

	for i := lo; i <= hi; i++ {
		left := generate(lo, i-1)
		right := generate(i+1, hi)

		for _, v := range left {
			for _, vv := range right {
				root := &TreeNode{Val: i}
				root.Left = v
				root.Right = vv
				ans = append(ans, root)
			}

		}
	}
	return ans
}
