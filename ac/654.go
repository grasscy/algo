package ac

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	cut := findMax(nums)
	root.Val = nums[cut]
	if cut-1 >= 0 {
		root.Left = constructMaximumBinaryTree(nums[:cut])
	} else {
		root.Left = nil
	}
	if cut+1 < len(nums) {
		root.Right = constructMaximumBinaryTree(nums[cut+1:])
	} else {
		root.Right = nil
	}
	return root
}

func findMax(nums []int) int {
	r := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[r] {
			r = i
		}
	}
	return r
}
