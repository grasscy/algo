package ac

func removeDuplicates(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	fast, slow := 1, 1
	pre := nums[0]
	for ; fast < len(nums); fast++ {
		if nums[fast] != pre {
			nums[slow] = nums[fast]
			slow++
		}
		pre = nums[fast]
	}
	return slow
}
