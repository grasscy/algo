package ac

func constructMaximumBinaryTree(nums []int) *TreeNode {
	r := build(nums, 0, len(nums))
	return r
}

func build(nums []int, start, end int) *TreeNode {
	if start >= end {
		return nil
	}

	max := nums[start]
	index := start
	for i := start; i < end; i++ {
		if max < nums[i] {
			max = nums[i]
			index = i
		}
	}

	root := &TreeNode{Val: max}
	root.Left = build(nums, start, index)
	root.Right = build(nums, index+1, end)
	return root
}
