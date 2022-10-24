package ac

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	index := findMax(nums)
	root := &TreeNode{Val: nums[index]}
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func findMax(nums []int) int {
	m := 0
	for i, v := range nums {
		if v > nums[m] {
			m = i
		}
	}
	return m
}
