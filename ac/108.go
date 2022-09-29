package ac

func sortedArrayToBST(nums []int) *TreeNode {
	return h(nums, 0, len(nums)-1)
}

func h(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}

	m := (l + r) / 2
	root := &TreeNode{nums[m], nil, nil}
	root.Left = h(nums, l, m-1)
	root.Right = h(nums, m+1, r)

	return root
}
